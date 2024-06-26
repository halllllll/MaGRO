import { createRootRouteWithContext, Outlet } from '@tanstack/react-router';
import { TanStackRouterDevtools } from '@tanstack/router-devtools';
import { Header } from '@/routes/-components/Header';
import { Box, Container, Flex, VStack, Text } from '@chakra-ui/react';
import { Footer } from '@/routes/-components/Footer';
import { LoginBtn } from '@/routes/-components/LoginButton';
import { AuthenticatedTemplate, UnauthenticatedTemplate } from '@azure/msal-react';
import { type QueryClient, QueryErrorResetBoundary } from '@tanstack/react-query';
import { Suspense } from 'react';
import { ErrorBoundary } from 'react-error-boundary';
import { ErrorFallback } from '@/components/ErrorFollback';
import { MsalProcess } from '@/components/MsalProcess';
import type { useEntraAuth } from '@/hooks/entraAuth';

interface RouterContext {
  azAuth: ReturnType<typeof useEntraAuth>;
  queryClient: QueryClient;
  unit: string | null;
}

export const Route = createRootRouteWithContext<RouterContext>()({
  component: () => (
    <Flex>
      <VStack
        height="100vh"
        overflow={{ base: 'auto', xl: 'hidden' }}
        _hover={{
          xl: { overflow: 'auto' },
        }}
        width={{ base: '100%', xl: 'calc(100% - 22rem)' }}
        padding="0 1rem"
      >
        <Container maxWidth="container.md" marginBottom="auto" px={'10'}>
          <Header />
          <Box mb={'auto'}>
            <MsalProcess />
            <AuthenticatedTemplate>
              <AuthenticatedTemplate>
                <QueryErrorResetBoundary>
                  {({ reset }) => (
                    <ErrorBoundary FallbackComponent={ErrorFallback} onReset={reset}>
                      <Suspense fallback={<b>Loading...</b>}>
                        <Outlet />
                      </Suspense>
                    </ErrorBoundary>
                  )}
                </QueryErrorResetBoundary>
              </AuthenticatedTemplate>
            </AuthenticatedTemplate>
            <UnauthenticatedTemplate>
              <Box>
                <Text as="i">
                  "MaGRO" is Maikurosohuto Graph-api account-password Reset Operator 🐟
                </Text>
              </Box>
              <LoginBtn />
            </UnauthenticatedTemplate>
          </Box>
        </Container>
        <Footer />
        <TanStackRouterDevtools />
      </VStack>
    </Flex>
  ),
});
