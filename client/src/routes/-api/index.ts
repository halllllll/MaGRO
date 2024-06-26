import { useSuspenseQuery } from '@tanstack/react-query';
import { getMaGROInfo } from './functions';
import { info } from './key';

export const useGetMaGROInfo = () => {
  const { data, isPending, isError, error } = useSuspenseQuery({
    staleTime: 10,
    gcTime: 30,
    queryFn: getMaGROInfo,
    queryKey: info.data,
  });

  if (isError) throw error;

  return { data, isPending, isError };
};
