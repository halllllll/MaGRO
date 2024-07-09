import { useMutation, useSuspenseQuery } from '@tanstack/react-query';
import type { UsersSubunitResponse, Auth, RepassRequest, RepassResponse } from './type';
import { belongsSubunitKeys } from './key';
import { getUnitData, repass } from './functions';
import type { GraphError } from '@/types/errors';

export const useGetUnitData = (authData: Auth, unitId: number) => {
  const { data, isPending, isError, error } = useSuspenseQuery<UsersSubunitResponse>({
    staleTime: 0,
    gcTime: 500,
    queryKey: belongsSubunitKeys.here(unitId),
    queryFn: () => getUnitData(unitId, authData.idToken),
  });

  return { data, isPending, isError, error };
};

export const useRepass = () => {
  const { mutate, isSuccess } = useMutation<RepassResponse, GraphError, RepassRequest>({
    mutationFn: repass,
    onMutate: (_req) => {
      // return req;
    },
    onSuccess: () => {},
  });

  return { mutate, isSuccess };
};
