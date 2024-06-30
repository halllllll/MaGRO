import { GetUnitID, RemoveUnitID } from '@/util/session';
import { Navigate, createFileRoute, redirect } from '@tanstack/react-router';

// いまんところ使われる予定はないので、どこかに確実にリダイレクトさせる
export const Route = createFileRoute('/unit/')({
  component: () => {
    return <Navigate to={'/'} />;
  },
  beforeLoad: () => {
    const storedUnit = GetUnitID();
    if (storedUnit) {
      throw redirect({
        to: '/unit/$unitId',
        params: {
          unitId: storedUnit,
        },
        replace: true,
        resetScroll: true,
      });
    }
    RemoveUnitID();
    throw redirect({
      to: '/',
      replace: true,
      resetScroll: true,
    });
  },
});
