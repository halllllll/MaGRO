import { RepeatIcon } from '@chakra-ui/icons';
import { Box, Flex, IconButton, Text } from '@chakra-ui/react';
import { Navigate, type ParsedLocation } from '@tanstack/react-router';
import type { FC } from 'react';

type Props = {
  loc: ParsedLocation;
};

export const Reflesh: FC<Props> = ({ loc }) => {
  return (
    <Box>
      <Flex gap={'3'} align={'center'}>
        <Text>情報を取得できませんでした。画面リロードを試してください</Text>
        <IconButton
          aria-label={'reload'}
          variant={'outline'}
          colorScheme={''}
          size={'md'}
          isRound={true}
          icon={<RepeatIcon />}
          mx={'2'}
          onClick={() => {
            Navigate({ to: loc.pathname, from: loc.pathname });
          }}
        />
      </Flex>
    </Box>
  );
};
