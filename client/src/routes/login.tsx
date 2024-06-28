import { Navigate, createFileRoute } from '@tanstack/react-router';
import { LoginBtn } from './-components/LoginButton';
import { Box, Text } from '@chakra-ui/react';
import type { FC } from 'react';
import { useEntraAuth } from '@/hooks/entraAuth';

const LoginComponent: FC = () => {
  const { isAuthenticated } = useEntraAuth();
  if (isAuthenticated) {
    console.info('yes loged in!!!');
    return <Navigate to="/" />;
  }
  return (
    <Box>
      <Box>
        <Text as="i">"MaGRO" is Maikurosohuto Graph-api account-password Reset Operator 🐟</Text>
      </Box>
      <LoginBtn />
    </Box>
  );
};

export const Route = createFileRoute('/login')({
  component: LoginComponent,
});
