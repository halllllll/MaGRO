import { useEntraAuth } from '@/hooks/entraAuth';
import { GetUnitID, RemoveUnitID, SetUnitID } from '@/util/session';
import { RepeatIcon } from '@chakra-ui/icons';
import { Box, Flex, IconButton, Text } from '@chakra-ui/react';
import {
  Navigate,
  createFileRoute,
  redirect,
  useLocation,
  useRouteContext,
} from '@tanstack/react-router';
import type { FC } from 'react';

const Component: FC = () => {
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
              Navigate({ to: '/' });
            }}
          />
        </Flex>
      </Box>
    );
  }

  // コンポーネント内でuseSuspenseQueryでデータ取得（loaderだとキャッシュが残らない)

  return <div>Hello /unit/$unitId!これちゃうん？</div>;
};

export const Route = createFileRoute('/unit/$unitId')({
  beforeLoad: ({ context: _context, params }) => {
    console.warn(`ほんとにとれてるの? /unitidだが... -> ${params.unitId}`);
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
    await acquireTokenSilent();
  },
  component: Component,
});
