import { yupResolver } from '@hookform/resolvers/yup';
import type { ReactNode } from '@tanstack/react-router';
import { useState, type FC } from 'react';
import { FormProvider, type SubmitHandler, useForm } from 'react-hook-form';
import { schema, type SchemaType } from './schema';
import { useDisclosure } from '@chakra-ui/react';
import { ModalOutline } from './Modal';
import type { SuccessData } from '../-api/type';
import type { User } from '@/entity/User';
import { UsersSubunitsList } from './Table';

export const Form: FC<{ children: ReactNode; data: SuccessData }> = ({ data }) => {
  // formのcheckboxで取っているのはuser idのみ。確認モーダルとかデータを表示するために、idからほかのユーザーデータを取得できるようにマップをつくる
  const candidateUser = new Map<string, User>();
  for (const u of data.user_groups) {
    candidateUser.set(u.user.user_id, u.user);
  }

  const { isOpen: isModalOpen, onOpen: onModalOpen, onClose: onModalClose } = useDisclosure();
  const methods = useForm<SchemaType>({
    mode: 'all',
    resolver: yupResolver(schema),
  });

  // TODO: 送信中の疑似挙動確認用
  const sleep = (ms: number) => new Promise((res) => setTimeout(res, ms));

  const [isProcessing, setIsProcessing] = useState<boolean>(false);

  const onSubmit: SubmitHandler<SchemaType> = async (data) => {
    onModalClose();
    console.log('pon!');
    console.log(data);
    setIsProcessing(true);
    await sleep(3000);
    setIsProcessing(false);
  };

  return (
    <>
      <FormProvider {...methods}>
        <ModalOutline
          isOpen={isModalOpen}
          onClose={onModalClose}
          data={candidateUser}
          onConfirm={onSubmit}
        />
        <form onSubmit={methods.handleSubmit(onModalOpen)}>
          <UsersSubunitsList data={data} isSending={isProcessing} />
        </form>
      </FormProvider>
    </>
  );
};
