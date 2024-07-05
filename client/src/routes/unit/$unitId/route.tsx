import { useEntraAuth } from '@/hooks/entraAuth';
import { GetUnitID, RemoveUnitID } from '@/util/session';
import { RepeatIcon } from '@chakra-ui/icons';
import { Box, Flex, IconButton, Text } from '@chakra-ui/react';
import {
  Navigate,
  createFileRoute,
  redirect,
  useLocation,
  useParams,
  useRouteContext,
} from '@tanstack/react-router';
import type { FC } from 'react';
import { useGetUnitData } from '../-api';
import { UsersSubunitsListX } from '../-comopnents/list';

const Component: FC = () => {
  const { IdToken, userId } = useEntraAuth();
  console.warn(`idtokenがほしいよ〜\n${IdToken}`);
  const unitId = useParams({ from: '/unit/$unitId', select: (params) => params.unitId });
  const ctx = useRouteContext({ from: '/unit/$unitId' });
  const loc = useLocation();
  console.log(`location? ${loc.pathname}`);
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
              Navigate({ to: loc.pathname, from: loc.pathname });
            }}
          />
        </Flex>
      </Box>
    );
  }

  // TODO: コンポーネント内でuseSuspenseQueryでデータ取得（loaderだとキャッシュが残らない)

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
      <div>Hello /unit/$unitId!これちゃうん？</div>
      <UsersSubunitsListX data={data.data} />
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
