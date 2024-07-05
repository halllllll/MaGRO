import { useSuspenseQuery } from '@tanstack/react-query';
import type { UsersSubunitResponse, Auth } from './type';
import { belongsSubunitKeys } from './key';
import { getUnitData } from './functions';

export const useGetUnitData = (authData: Auth, unitId: number) => {
  const { data, isPending, isError, error } = useSuspenseQuery<UsersSubunitResponse>({
    staleTime: 0,
    gcTime: 500,
    queryKey: belongsSubunitKeys.here(unitId),
    queryFn: () => getUnitData(unitId, authData.idToken),
  });

  return { data, isPending, isError, error };
};
