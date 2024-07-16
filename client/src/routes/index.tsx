import {
  createFileRoute,
  Link,
  redirect,
  useNavigate,
  useRouteContext,
} from '@tanstack/react-router';
import { Box, Flex, IconButton, ListItem, Text, UnorderedList } from '@chakra-ui/react';
import { Suspense, type FC } from 'react';
import { useGetBelongingUnits } from './-api';
import { QueryErrorResetBoundary } from '@tanstack/react-query';
import { useEntraAuth } from '@/hooks/entraAuth';
import { ErrorFallback } from '@/components/ErrorFollback';
import { ErrorBoundary } from 'react-error-boundary';
import { GetUnitID, SetUnitID } from '@/util/session';
import { RepeatIcon } from '@chakra-ui/icons';

const Component: FC = () => {
  const { IdToken, userId, acquireTokenSilent } = useEntraAuth();
  const navigate = useNavigate({ from: '/login' });
  // TODO: エラーハンドリング
  const { data } = useGetBelongingUnits({
    userId: userId,
    idToken: IdToken,
  });

  // contextがほしいがなぜかとれない（画面リロードで取れる）ことがある
  const ctx = useRouteContext({ from: '/' });
  if (!ctx) {
    return (
      <Box>
        <Flex gap={'3'} align={'center'}>
          <Text>情報を取得できませんでした。画面リロードを試してください</Text>
          <IconButton
            aria-label={'reload'}
            variant={'outline'}
            colorScheme={''}
            size={'md'}
            isRound={true}
            icon={<RepeatIcon />}
            mx={'2'}
            onClick={() => {
              navigate({ to: '/' });
            }}
          />
        </Flex>
      </Box>
    );
  }

  if (data?.status === 'error') {
    throw new Error(data.message);
  }
  // TODO: 選択肢がひとつの場合はそのままリダイレクトしたいが、
  // なぜかエラーが解決できないので、とりあえず単一であっても候補を選ばせることにする
  /*
  // if (data?.unit_count === 1) {
  //   // TODO: 諦めてリダイレクト先でもっかいリクエストを飛ばすことにしている
  //   // (データの渡し方がわからないし存在するのかも不明)
  //   SetUnitID(data.units[0].unit_id);
  //   // cotextの更新の仕方がわからない
  //   // ctx.unit[1](GetUnitID()); なくした
  //   // componentの中だとthrow redirectじゃなくてnavidate?
  //   navigate({
  //     to: '/unit/$unitId',
  //     params: {
  //       unitId: data.units[0].unit_id,
  //     },
  //     replace: true,
  //     resetScroll: true,
  //   });
  //   // throwを使うとuncauht errorなる
  //   // throw redirect({
  //   //   to: '/unit/$unitId',
  //   //   params: {
  //   //     unitId: data.units[0].unit_id,
  //   //   },
  //   //   replace: true,
  //   //   resetScroll: true,
  //   // });
  // }
  */

  const onResetError = async () => {
    // TODO: 意味があるのかどうかわからない。
    // -> Auth Errorのときは再ログインさせたい
    await acquireTokenSilent();
  };

  const { unit } = ctx;
  const [_, setUnitName] = unit;

  return (
    <>
      <QueryErrorResetBoundary>
        {({ reset: _reset }) => (
          <ErrorBoundary onReset={onResetError} FallbackComponent={ErrorFallback}>
            <Suspense fallback={<h2>fetching...</h2>}>
              {data.unit_count === 0 ? (
                <div>登録されているUnitはありません</div>
              ) : (
                <>
                  <Box>
                    <Text>ログイン先を選んでください</Text>
                    <UnorderedList>
                      {data.units.map((v) => {
                        return (
                          <Link
                            to="/unit/$unitId"
                            params={{
                              unitId: v.unit_id.toString(),
                            }}
                            key={v.unit_id}
                          >
                            <ListItem
                              key={v.unit_id}
                              onClick={() => {
                                setUnitName(v.name);
                                SetUnitID(v.unit_id.toString());
                              }}
                            >
                              {v.name}
                            </ListItem>
                          </Link>
                        );
                      })}
                    </UnorderedList>
                  </Box>
                </>
              )}
            </Suspense>
          </ErrorBoundary>
        )}
      </QueryErrorResetBoundary>
    </>
  );
};

export const Route = createFileRoute('/')({
  component: Component,
  // TODO: なぜかcontext is undefinedみたいなエラーが出る（ここが原因かは不明）
  // https://github.com/TanStack/router/issues/1531
  // https://github.com/TanStack/router/issues/1751
  beforeLoad: async ({ context }) => {
    if (!context) {
      return;
    }
    // なぜかloaderではcontextがないと言われてしまうのでこっちにもってきた
    const { acquireTokenSilent } = context.azAuth;
    await acquireTokenSilent();
    // すでにunitが選択された状態(session storage)だったらそこに飛ばす
    const storedUnit = GetUnitID();
    if (storedUnit) {
      throw redirect({
        to: '/unit/$unitId',
        params: {
          unitId: storedUnit,
        },
        replace: true,
        resetScroll: true,
      });
    }
  },

  loader: async ({ context: _context }) => {
    // 画面に関わることなのでcomponentにしようと思ったが、結果が一つの場合で分けたいので
    // なぜかcontextがないと言われてしまうのでbeforeloadに移動した
    // TODO: サーバーでもやっているが、ロール確認のデータをここで取得したい(cacheする必要のないデータ)
    // componentではuseLoaderDataで取得できる
  },

  pendingComponent: () => {
    return <>{'waiting...'}</>;
  },
  notFoundComponent: () => {
    // TODO: なぜか意味がない
    throw redirect({
      to: '/',
      replace: true,
      resetScroll: true,
      viewTransition: false,
    });
  },
});
