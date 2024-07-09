import { useQuery, useSuspenseQuery } from '@tanstack/react-query';
import { getMaGROInfo, getUnitsByUser } from './functions';
import { belongsUnitKeys, info } from './key';
import type { BelongUnitsResponse, Auth } from './type';

export const useGetMaGROInfo = (authData: Auth) => {
  const { data, isPending, isError, error } = useQuery({
    staleTime: 0,
    gcTime: 300,
    queryFn: () => getMaGROInfo(authData.idToken),
    queryKey: info.me(authData.userId),
  });

  if (isError) throw error;

  return { data, isPending, isError };
};

export const useGetBelongingUnits = (authData: Auth) => {
  const { data, isPending, isError, error } = useSuspenseQuery<BelongUnitsResponse /*APIError*/>({
    staleTime: 0,
    gcTime: 300,
    queryFn: () => getUnitsByUser(authData.idToken),
    queryKey: belongsUnitKeys.me(authData.userId),
  });

  // if (isError) throw error;

  return { data, isPending, isError, error };
};
