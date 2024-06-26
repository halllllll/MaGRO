import { createFileRoute } from '@tanstack/react-router';

export const Route = createFileRoute('/redirect')({
  component: () => <div>Hello /redirect!</div>,
});
