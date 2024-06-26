import { Box, Text } from '@chakra-ui/react';
import type { FC } from 'react';

export const Footer: FC = () => {
  return (
    <Box>
      <footer style={{ marginTop: '1rem', width: '100%' }}>
        <Text padding="1rem" textAlign="center" fontSize=".8rem">
          &copy; 2024 GIGSCHOOL
        </Text>
      </footer>
    </Box>
  );
};
