import { useEntraAuth } from '@/hooks/entraAuth';
import { routeTree } from '@/routeTree.gen';
import { RouterProvider, createRouter } from '@tanstack/react-router';
import { queryClient } from './QueryProvider';

// new tanstack router instance
const router = createRouter({
  routeTree,
  // defaultPreload: 'intent',
});

declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router;
  }
}

export const MaGRORouterProvider = () => {
  // boforeloaderとかで使うhookはここで宣言しといてcontextで渡す(createRootRouteWithContext)
  const ctx = {
    azAuth: useEntraAuth(),
    queryClient: queryClient,
    // unit: useState<string | null>(null),
  };

  return <RouterProvider router={router} context={ctx} />;
};
