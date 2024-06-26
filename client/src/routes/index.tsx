import { createFileRoute, Link } from '@tanstack/react-router';
import { Box } from '@chakra-ui/react';
import type { FC } from 'react';

const Component: FC = () => {
  return (
    <>
      <Box>yes?</Box>
      <Link to="/user">{'>aaa<'}</Link>
    </>
  );
};

export const Route = createFileRoute('/')({
  component: Component,
  loader: async ({ context }) => {
    // routerとqueryのテスト
    // 画面に関わることなのでcomponentにしようと思ったが、結果が一つの場合で分けたいので
    // get user msal data
    const { userId, IdToken } = context.azAuth;

    // query client
    const queryClient = context.queryClient;
    const data = await queryClient.fetchQuery({
      queryKey: [userId],
      queryFn: async () => {
        const res = await fetch('/api/info', {
          method: 'GET',
          headers: {
            Authorization: `Bearer ${IdToken}`,
          },
        });

        if (!res.ok) {
          const err = await res.json();
          throw new Error(err);
        }
        return await res.json();
      },
    });
    console.log(data);
  },
  pendingComponent: () => {
    return <>{'waiting...'}</>;
  },
});
