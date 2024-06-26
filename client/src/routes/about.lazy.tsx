import { Link, createLazyFileRoute } from '@tanstack/react-router';

export const Route = createLazyFileRoute('/about')({
  component: About,
});

function About() {
  return (
    <>
      <div className="p-2">Hello from About!</div>
      <div>aaa?</div>
      <Link to={'/select_unit'}>aaaaa</Link>
    </>
  );
}
