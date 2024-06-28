import { createFileRoute, Link, redirect } from '@tanstack/react-router';
import { Box } from '@chakra-ui/react';
import { Suspense, type FC } from 'react';
import { useGetMaGROInfo } from './-api';
import { QueryErrorResetBoundary } from '@tanstack/react-query';
import { useEntraAuth } from '@/hooks/entraAuth';
import { ErrorFallback } from '@/components/ErrorFollback';
import { ErrorBoundary } from 'react-error-boundary';

const Component: FC = () => {
  const { IdToken, userId, acquireTokenSilent } = useEntraAuth();
  const { data } = useGetMaGROInfo({
    userId: userId,
    idToken: IdToken,
  });

  console.log(`data! ${data}`);

  const onResetError = async () => {
    // TODO: 意味があるのかどうかわからない。
    // -> Auth Errornのときは再ログインさせたい
    console.log('reflesh!');
    await acquireTokenSilent();
  };

  return (
    <>
      <QueryErrorResetBoundary>
        {({ reset: _reset }) => (
          <ErrorBoundary onReset={onResetError} FallbackComponent={ErrorFallback}>
            <Suspense fallback={<h2>fetching...</h2>}>
              <Box>yes?</Box>
              <Link to="/user">{'>aaa<'}</Link>
            </Suspense>
          </ErrorBoundary>
        )}
      </QueryErrorResetBoundary>
    </>
  );
};

export const Route = createFileRoute('/')({
  component: Component,
  // TODO: なぜかcontext is undefinedみたいなエラーが出る（ここが原因化かどうか不明）
  // beforeLoad: ({ context }) => {
  //   if (context === undefined) {
  //     console.warn('????');
  //     redirect({
  //       to: '/login',
  //       replace: true,
  //     });
  //   }
  //   // もしすでにunitが選択された状態だったらそこに飛ばす
  //   const storedUnit = sessionStorage.getItem('unit_id');
  //   if (storedUnit) {
  //     context.unit = storedUnit;
  //     redirect({
  //       to: `/unit/${storedUnit}`,
  //       replace: true,
  //       resetScroll: true,
  //     });
  //   }
  // },

  loader: async ({ context }) => {
    // routerとqueryのテスト
    // 画面に関わることなのでcomponentにしようと思ったが、結果が一つの場合で分けたいので
    // get user msal data
    const { acquireTokenSilent } = context.azAuth;
    console.warn('uouo~~~');
    await acquireTokenSilent();
  },

  pendingComponent: () => {
    return <>{'waiting...'}</>;
  },
  notFoundComponent: () => {
    // TODO: なぜか意味がない
    throw redirect({
      to: '/',
      replace: false,
    });
  },
});
