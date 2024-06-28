import {
  type AccountInfo,
  InteractionStatus,
  BrowserAuthError,
  InteractionRequiredAuthError,
  type SilentRequest,
  EventType,
} from '@azure/msal-browser';
import { useMsal, useIsAuthenticated } from '@azure/msal-react';
import { useMemo, useCallback } from 'react';
import { AppRequests } from './graphScope';

let tokenExpirationTimer: ReturnType<typeof setTimeout> | undefined;

export const useEntraAuth = () => {
  const { instance, inProgress } = useMsal();

  const isAuthenticated = useIsAuthenticated();

  const userId = useMemo(() => {
    const accounts = instance.getAllAccounts();
    if (accounts.length > 0) {
      const account: AccountInfo = accounts[0];
      if (account?.idTokenClaims) {
        return account.idTokenClaims.sub || '';
      }
      return '';
    }
  }, [instance]);

  // retrieve silent token
  // ex: (from invoke GET ACCESS TOKEN): const accessToken = await msAuth.acquireTokenSilent();
  const acquireTokenSilent = useCallback(async (): Promise<string | null> => {
    // TODO: AuthErrorになったときの挙動確認（IdTokenがExpireされてたら再ログインを促したい）
    console.warn('更新にチャレンジ！');
    const accounts = instance.getAllAccounts();
    if (accounts.length > 0) {
      const account = accounts[0];
      instance.setActiveAccount(account);
      try {
        const response = await instance.acquireTokenSilent({
          account,
          forceRefresh: true,
          ...AppRequests,
        });
        return response.accessToken;
      } catch (err) {
        console.error('Error acquiring token silently:', err);
        return null;
      }
    } else {
      console.error('No accounts found');
      return null;
    }
  }, [instance]);

  // TODO: IdTokenは更新されない 再ログインを促す？
  const setupTokenExpirationTimer = (): void => {
    const accounts = instance.getAllAccounts();
    if (accounts.length > 0) {
      const account = accounts[0];
      const exp = account.idTokenClaims?.exp;
      if (typeof exp === 'number') {
        // トークンの有効期限までの時間を計算
        const tokenExpirationTime = exp * 1000;
        const currentTime = Date.now();
        const timeUntilExpiration = tokenExpirationTime - currentTime;

        // 古いタイマーをクリアして新しいタイマーを設定
        if (tokenExpirationTimer) clearTimeout(tokenExpirationTimer);
        tokenExpirationTimer = setTimeout(() => {
          refreshAccessToken(account);
        }, timeUntilExpiration);
      }
    }
  };

  // biome-ignore lint/correctness/useExhaustiveDependencies: <explanation>
  const loginAzure = useCallback(async () => {
    if (inProgress === InteractionStatus.None) {
      // instance.handleRedirectPromise()
      await instance.loginRedirect(AppRequests);
      // 成功時のコールバック
      instance.addEventCallback((eve) => {
        if (eve.eventType === EventType.LOGIN_SUCCESS) {
          setupTokenExpirationTimer();
        }
      });
    }
  }, [instance]);

  // TODO: will back when get userhomeid and replace ↑
  // const logoutAzure = useCallback(
  //   async (accountId: string) => {
  //     const targetAccount = instance.getAccountByHomeId(accountId);
  //     if (targetAccount) {
  //       await instance.logoutRedirect({ account: targetAccount });
  //     } else {
  //       console.error('account not found');
  //     }
  //   },
  //   [instance],
  // );

  const logoutAzure = useCallback(async () => {
    instance.logoutRedirect();
  }, [instance]);

  const refreshAccessToken = async (account: AccountInfo): Promise<void> => {
    const silentRequest: SilentRequest = {
      account,
      ...AppRequests,
    };

    try {
      // try silent reflesh
      const response = await instance.acquireTokenSilent(silentRequest);
      console.log('Refreshed Access Token:', response.accessToken);
      setupTokenExpirationTimer();
    } catch (err) {
      // InteractionRequiredAuthError / BrowserAuthError エラーの場合、再度リダイレクトで認証させる
      if (err instanceof InteractionRequiredAuthError) {
        console.info('redirect...');
        await instance.acquireTokenRedirect(AppRequests);
        setupTokenExpirationTimer();
      } else if (err instanceof BrowserAuthError) {
        console.info('redirect...');
        await instance.acquireTokenRedirect(AppRequests);
        setupTokenExpirationTimer();
      } else {
        console.error('Error refreshing access token:', err);
      }
    }
  };

  const IdToken = useMemo(() => {
    const accounts = instance.getAllAccounts();
    if (accounts.length > 0) {
      const account: AccountInfo = accounts[0];
      if (account?.idToken) {
        return account.idToken;
      }
      return '';
    }
  }, [instance]);

  return {
    acquireTokenSilent,
    loginAzure,
    logoutAzure,
    inProgress,
    isAuthenticated,
    instance,
    IdToken,
    userId, // TODO: 不要？
  };
};
