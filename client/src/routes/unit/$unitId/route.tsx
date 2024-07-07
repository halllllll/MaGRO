import { useEntraAuth } from '@/hooks/entraAuth';
import { GetUnitID, RemoveUnitID } from '@/util/session';
import {
  createFileRoute,
  redirect,
  useLocation,
  useParams,
  useRouteContext,
} from '@tanstack/react-router';
import { Suspense, type FC } from 'react';
import { useGetUnitData } from '../-api';
import { Reflesh } from '../-comopnents/reflesh';
import { Form } from '../-comopnents/Form';
import { QueryErrorResetBoundary } from '@tanstack/react-query';
import { ErrorBoundary } from 'react-error-boundary';
import { ErrorFallback } from '@/components/ErrorFollback';
import { PageInfo } from '../-comopnents/pageInfo';

const Component: FC = () => {
  const { IdToken, userId } = useEntraAuth();
  console.warn(`idtokenがほしいよ〜\n${IdToken}`);
  const unitId = useParams({ from: '/unit/$unitId', select: (params) => params.unitId });
  const ctx = useRouteContext({ from: '/unit/$unitId' });

  const loc = useLocation();
  console.log(`location? ${loc.pathname}`);
  if (!ctx) {
    <Reflesh loc={loc} />;
  }

  // コンポーネント内でuseSuspenseQueryでデータ取得（loaderだとキャッシュが残らない)
  const { data } = useGetUnitData({ userId: userId, idToken: IdToken }, Number.parseInt(unitId));
  console.warn('data!');
  console.dir(data);
  if (data.status === 'error') {
    return (
      <>
        {data.status} {data.message}
      </>
    );
  }

  return (
    <>
      <QueryErrorResetBoundary>
        {({ reset }) => (
          <ErrorBoundary FallbackComponent={ErrorFallback} onReset={reset}>
            <Suspense fallback={<b>loading...</b>}>
              <PageInfo />
              <Form data={data.data}>{''}</Form>
            </Suspense>
          </ErrorBoundary>
        )}
      </QueryErrorResetBoundary>
    </>
  );
};

export const Route = createFileRoute('/unit/$unitId')({
  beforeLoad: ({ context: _context, params }) => {
    console.warn(`unitはcontext併用ではなくsessiondだけで管理することにした ${GetUnitID()}`);
    const storedUnit = GetUnitID();
    console.warn(`saved unit id? ${storedUnit}`);
    if (!storedUnit) {
      RemoveUnitID();
      throw redirect({
        to: '/',
        params: {
          unitId: storedUnit,
        },
        replace: true,
        resetScroll: true,
      });
    }
    if (storedUnit !== params.unitId) {
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
  loader: async ({ context }) => {
    const { acquireTokenSilent } = context.azAuth;
    const at = await acquireTokenSilent();
    console.warn(`access tokenだよ〜 ${at}`);
  },
  component: Component,
  // gcTime: 0, // TODO: for dev
  // staleTime: 0, // TODO: for dev
});
