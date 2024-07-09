import { Box, Divider, Flex, Heading, Spacer, Text, useMediaQuery } from '@chakra-ui/react';
import type { FC } from 'react';
import { Link, useRouteContext } from '@tanstack/react-router';
import { AuthenticatedTemplate, useMsal } from '@azure/msal-react';
import { LogoutBtn } from './LogoutButton';

export const Header: FC = () => {
  const [isTablet] = useMediaQuery('(min-width: 48em)');

  return (
    <>
      <Heading position={isTablet ? 'relative' : 'sticky'} width={'100%'}>
        <Flex
          align="center"
          justify="center"
          boxSize="full"
          height="16"
          width={'100%'}
          gap={'1.5rem'}
        >
          <Box>
            <Link to="/">MaGRO</Link>
          </Box>
          <Spacer />
          <AuthenticatedTemplate>
            <Flex direction={'column'} gap={2}>
              <UnitSan />
              <AccountSan />
            </Flex>
            <LogoutBtn />
          </AuthenticatedTemplate>
        </Flex>
        <Divider />
      </Heading>
    </>
  );
};

const AccountSan = () => {
  const { accounts } = useMsal();
  const account = accounts[0];
  return <Text fontSize={'14'}>{`${account.name}`}</Text>;
};

// TODO: なんかuseRouteContextでエラーになる
const UnitSan = () => {
  try {
    const ctx = useRouteContext({ from: '/unit/' });
    if (!ctx) return null;
    const { unit } = ctx;
    const [unitName, _] = unit;
    return unitName ? <Text fontSize={'16'}>{unitName}</Text> : null;
  } catch {
    return null;
  }
};
