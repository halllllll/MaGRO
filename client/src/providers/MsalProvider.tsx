import type { FC, ReactNode } from 'react';

import { BrowserCacheLocation, LogLevel, PublicClientApplication } from '@azure/msal-browser';
import { MsalProvider } from '@azure/msal-react';

// BowtherAuthOptions
type AuthProperties = {
  auth: {
    clientId: string;
    authority: string;
    redirectUri: string;
    postLogoutRedirectUri?: string | null;
    navigateToLoginRequestUrl?: boolean;
  };
};

export const createMsalClient = (auth: AuthProperties) => {
  const msalClient = new PublicClientApplication({
    ...auth,
    cache: {
      cacheLocation: BrowserCacheLocation.SessionStorage,
      storeAuthStateInCookie: true,
    },
    system: {
      loggerOptions: {
        logLevel: LogLevel.Trace,
        loggerCallback: (level, message, containsPii) => {
          if (containsPii) {
            return;
          }
          switch (level) {
            case LogLevel.Error:
              console.error(message);
              return;
            case LogLevel.Info:
              console.info(message);
              return;
            case LogLevel.Verbose:
              console.debug(message);
              return;
            case LogLevel.Warning:
              console.warn(message);
              return;
            default:
              console.log(message);
              return;
          }
        },
      },
      tokenRenewalOffsetSeconds: 300,
    },
  });

  return msalClient;
};

export const MsalClientProvider: FC<{ children: ReactNode }> = ({ children }) => {
  const msalClient = createMsalClient({
    auth: {
      clientId: import.meta.env.VITE_CLIENT_ID ?? '',
      authority: `https://login.microsoftonline.com/${import.meta.env.VITE_AUTHORITY}`,
      redirectUri: import.meta.env.DEV
        ? `http://localhost:${import.meta.env.VITE_PORT}/${import.meta.env.VITE_REDIRECT_URI}`
        : `${import.meta.env.VITE_URI}/${import.meta.env.VITE_REDIRECT_URI}`,
      navigateToLoginRequestUrl: false,
      postLogoutRedirectUri: '/',
    },
  });
  return <MsalProvider instance={msalClient}>{children}</MsalProvider>;
};
