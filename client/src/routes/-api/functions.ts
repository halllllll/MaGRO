import type { Info } from './type';
import { useAzureAuth } from '@/hooks/entraAuth';

export const getMaGROInfo = async (): Promise<Info> => {
  const { IdToken } = useAzureAuth();
  const res = await fetch('/api/info', {
    method: 'GET',
    headers: {
      Authorization: `Bearer ${IdToken}`,
    },
  });

  if (!res.ok) {
    throw new Error(`${res.status} ${res.statusText}`);
  }

  return await res.json();
};
