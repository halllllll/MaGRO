import { createFileRoute, redirect } from '@tanstack/react-router';

export const Route = createFileRoute('/user/')({
  component: () => Component(),
  beforeLoad: ({ context }) => {
    const { isAuthenticated } = context.azAuth;
    if (!isAuthenticated) {
      throw redirect({
        to: '/',
        replace: true,
      });
    }
  },
  errorComponent: () => {
    <h1>Error</h1>;
  },
});

const Component = () => {
  return (
    <>
      <div>Hello /user/!</div>
      <div>yes!</div>
    </>
  );
};
