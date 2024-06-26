import { ChakraProvider } from '@chakra-ui/react';
import { theme } from '../theme';

import { MsalClientProvider } from './MsalProvider';
import { QueryProvider } from './QueryProvider';
import { MaGRORouterProvider } from './MaGRORouterProvider';

export const Providers = () => {
  return (
    <MsalClientProvider>
      <QueryProvider>
        <ChakraProvider theme={theme}>
          <MaGRORouterProvider />
        </ChakraProvider>
      </QueryProvider>
    </MsalClientProvider>
  );
};
