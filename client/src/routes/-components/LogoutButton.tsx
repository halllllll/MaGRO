import { useAzureAuth } from '@/hooks/entraAuth';
import { InteractionStatus } from '@azure/msal-browser';
import { Button } from '@chakra-ui/react';
import type { FC } from 'react';

export const LogoutBtn: FC = () => {
  const { logoutAzure, inProgress } = useAzureAuth();
  return (
    <Button
      isLoading={InteractionStatus.Login === inProgress}
      disabled={InteractionStatus.Login === inProgress}
      onClick={logoutAzure}
    >
      Logout
    </Button>
  );
};
