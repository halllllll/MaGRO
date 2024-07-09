import { yupResolver } from '@hookform/resolvers/yup';
import { useParams, type ReactNode } from '@tanstack/react-router';
import { useState, type FC } from 'react';
import { FormProvider, type SubmitHandler, useForm } from 'react-hook-form';
import { schema, type SchemaType } from './schema';
import { useDisclosure, useToast } from '@chakra-ui/react';
import { ConfirmModal } from './ConfirmModal';
import type { RepassResultData, RepassRequest, SuccessData, TargetUsers } from '../-api/type';
import type { User } from '@/entity/User';
import { UsersSubunitsList } from './Table';
import { useRepass } from '../-api';
import { useEntraAuth } from '@/hooks/entraAuth';
import { ResultModal } from './ResultModal';

export const Form: FC<{ children: ReactNode; data: SuccessData }> = ({ data }) => {
  const currentUser = data.current_user;
  // formのcheckboxで取っているのはuser idのみ。確認モーダルとかデータを表示するために、idからほかのユーザーデータを取得できるようにマップをつくる
  const candidateUser = new Map<string, User>();
  for (const u of data.user_groups) {
    candidateUser.set(u.user.user_id, u.user);
  }

  const toast = useToast();
  // 成功時はフルモーダルで情報を出すことにした（n敗　ホントは別画面に遷移したいがデータをそっちに移す方法がわからない）ので、それを取得するstate
  const [resultData, setResultData] = useState<RepassResultData[]>([]);

  const {
    isOpen: isConfirmModalOpen,
    onOpen: onConfirmModalOpen,
    onClose: onConfirmModalClose,
  } = useDisclosure();
  const {
    isOpen: isResultModalOpen,
    onOpen: onResultModalOpen,
    onClose: onResultModalClose,
  } = useDisclosure();

  const methods = useForm<SchemaType>({
    mode: 'all',
    resolver: yupResolver(schema),
  });

  const { IdToken, userId } = useEntraAuth();
  const unitId = useParams({ from: '/unit/$unitId', select: (params) => params.unitId });

  const { mutate, isSuccess } = useRepass();

  // TODO: 送信中の疑似挙動確認用
  // const sleep = (ms: number) => new Promise((res) => setTimeout(res, ms));

  const [isProcessing, setIsProcessing] = useState<boolean>(false);

  const onSubmit: SubmitHandler<SchemaType> = async (data) => {
    onConfirmModalClose();
    // 送信用のデータ構築
    const userpayload: TargetUsers[] = data.user_ids
      .map((v) => candidateUser.get(v ?? '') ?? '')
      .filter((v) => v !== '') // なんかfilterをまとめられなかった(`''`がなぜか残る扱い)が、そこまで計算量多くないのでいったん無視する
      .filter((v, i, s) => s.indexOf(v) === i) // 重複削除
      .map((v) => {
        return {
          user_id: v.user_id,
          user_account: v.user_name,
        };
      });
    const reqData: RepassRequest = {
      auth: { userId: userId, idToken: IdToken },
      unitId: Number.parseInt(unitId),
      target_user: userpayload,
      current_user: currentUser,
    };

    mutate(
      { ...reqData },
      {
        onSettled: () => {},
        onSuccess: (data) => {
          if (data.status === 'success') {
            methods.reset();
            setIsProcessing(true);
            setResultData(data.body);
            onResultModalOpen();
          } else {
            throw new Error(data.message);
          }
        },
        onError: (error) => {
          console.warn(error);
          toast({
            title: '失敗しました',
            description: `${error.name} - ${error.message}`,
            status: 'error',
            isClosable: true,
            duration: 8000,
          });
        },
      },
    );
  };

  return (
    <>
      <FormProvider {...methods}>
        <ConfirmModal
          isOpen={isConfirmModalOpen}
          onClose={onConfirmModalClose}
          data={candidateUser}
          onConfirm={onSubmit}
        />
        {isSuccess && (
          <ResultModal
            isOpen={isResultModalOpen}
            onClose={() => {
              onResultModalClose();
              setIsProcessing(false);
            }}
            data={resultData}
          />
        )}
        <form onSubmit={methods.handleSubmit(onConfirmModalOpen)}>
          <UsersSubunitsList data={data} isSending={isProcessing} />
        </form>
      </FormProvider>
    </>
  );
};
