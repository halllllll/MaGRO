import { createFileRoute } from '@tanstack/react-router';

export const Route = createFileRoute('/select_unit')({
  component: () => <div>Hello /select_unit!</div>,
});
