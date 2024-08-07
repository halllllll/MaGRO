import { useEntraAuth } from '@/hooks/entraAuth';
import { InteractionStatus } from '@azure/msal-browser';
import { Button } from '@chakra-ui/react';
import type { FC } from 'react';

export const LogoutBtn: FC = () => {
  const { logoutAzure, inProgress } = useEntraAuth();
  return (
    <Button
      isLoading={InteractionStatus.Login === inProgress}
      disabled={InteractionStatus.Login === inProgress}
      variant={"outline"}
      onClick={logoutAzure}
    >
      Logout
    </Button>
  );
};
