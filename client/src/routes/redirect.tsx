import { createFileRoute, Navigate } from '@tanstack/react-router';

export const Route = createFileRoute('/redirect')({
  component: () => <div>yay!</div>, // <Navigate to="/" />,
});
