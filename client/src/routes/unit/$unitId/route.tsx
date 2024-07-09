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
  const unitId = useParams({ from: '/unit/$unitId', select: (params) => params.unitId });
  const ctx = useRouteContext({ from: '/unit/$unitId' });

  const loc = useLocation();
  if (!ctx) {
    <Reflesh loc={loc} />;
  }

  // コンポーネント内でuseSuspenseQueryでデータ取得（loaderだとキャッシュが残らない)
  const { data } = useGetUnitData({ userId: userId, idToken: IdToken }, Number.parseInt(unitId));
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
    // TODO: unitはcontext併用ではなくsessiondだけで管理することにした
    const storedUnit = GetUnitID();

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
    // * if want get accesstoken, ↓
    // const accessToken = await acquireTokenSilent();
    await acquireTokenSilent();
  },
  component: Component,
  // gcTime: 0, // TODO: for dev
  // staleTime: 0, // TODO: for dev
});
