import { useAzureAuth } from '@/hooks/entraAuth';
import { InteractionStatus } from '@azure/msal-browser';
import { Box, Flex, Spinner, Text } from '@chakra-ui/react';
import { type FC, useMemo } from 'react';

export const MsalProcess: FC = () => {
  const { inProgress } = useAzureAuth();
  const message = useMemo(() => {
    switch (inProgress) {
      case InteractionStatus.Startup:
        return 'Setting things up';
      case InteractionStatus.Login:
        return 'Logging you in';
      case InteractionStatus.Logout:
        return 'Logging you out';
      case InteractionStatus.HandleRedirect:
        return 'Redirecting...';
      case InteractionStatus.AcquireToken:
        return 'Acquire token...';
      case InteractionStatus.None:
        return ' ';
      default:
        return ' ';
    }
  }, [inProgress]);
  return (
    <Box minH={8}>
      {message !== ' ' && (
        <Flex gap={5}>
          <Spinner />
          <Text>{`${message}`}</Text>
        </Flex>
      )}
    </Box>
  );
};
