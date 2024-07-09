import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  Box,
  Button,
  Text,
  ModalBody,
  ModalCloseButton,
  UnorderedList,
  ModalFooter,
  VStack,
  Alert,
  AlertDescription,
  AlertIcon,
  AlertTitle,
  ListItem,
  TableContainer,
  Tbody,
  Tr,
  Table,
  Td,
  Thead,
  TableCaption,
  Th,
  Divider,
} from '@chakra-ui/react';
import type { FC } from 'react';
import type { RepassResultData } from '../-api/type';

type Props = {
  isOpen: boolean;
  onClose: () => void;
  data: RepassResultData[];
};

export const ResultModal: FC<Props> = (props) => {
  const { isOpen, onClose, data } = props;

  const failedUser = data.filter((v) => v.status === 'error');
  const successedUser = data.filter((v) => v.status === 'success');

  return (
    <Modal
      isOpen={isOpen}
      size={'full'}
      onClose={onClose}
      closeOnOverlayClick={false}
      allowPinchZoom={true}
      scrollBehavior={'inside'}
    >
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>MaGRO Result</ModalHeader>
        <>
          <ModalCloseButton />
          <ModalBody pb={6}>
            <Box>
              <Alert status="error">
                <AlertIcon />
                <AlertTitle>{''}</AlertTitle>
                <Box>
                  <Text as={'b'}>ATTENTION</Text>
                  <UnorderedList styleType="'- '">
                    <ListItem>
                      <AlertDescription>この画面は一度しか表示されません</AlertDescription>
                    </ListItem>
                  </UnorderedList>
                </Box>
              </Alert>
            </Box>
            <Box>
              <TableContainer w={'100%'} whiteSpace={'unset'} overflowX="unset" overflowY="unset">
                {failedUser.length > 0 && (
                  <Table variant="simple" colorScheme={'yellow.200'} backgroundColor={'yellow.100'}>
                    <TableCaption placement={'top'}>Failed</TableCaption>
                    <Thead>
                      <Tr>
                        <Th>ID</Th>
                        <Th>Name</Th>
                        <Th>Message</Th>
                      </Tr>
                    </Thead>
                    <Tbody>
                      {failedUser.map((v, i) => {
                        return (
                          <Tr key={`${i}_${v.user.user_id}`}>
                            <Td>{v.user.user_name}</Td>
                            <Td>{v.user.user_displayname}</Td>
                            <Td>{v.message}</Td>
                          </Tr>
                        );
                      })}
                    </Tbody>
                  </Table>
                )}

                <Divider size={'10'} />
                {successedUser.length > 0 && (
                  <Table variant="simple">
                    <TableCaption placement={'top'}>Success</TableCaption>
                    <Thead>
                      <Tr>
                        <Th>ID</Th>
                        <Th>Name</Th>
                        <Th>New Password</Th>
                      </Tr>
                    </Thead>
                    <Tbody>
                      {successedUser.map((v, i) => {
                        return (
                          <Tr key={`${i}_${v.user.user_id}`}>
                            <Td>{v.user.user_name}</Td>
                            <Td>{v.user.user_displayname}</Td>
                            <Td>
                              <Text as="b">{`${v.password}`}</Text>
                            </Td>
                          </Tr>
                        );
                      })}
                    </Tbody>
                  </Table>
                )}
              </TableContainer>
            </Box>
            <VStack spacing={4} />
          </ModalBody>
          <ModalFooter>
            <Button onClick={onClose}>Close</Button>
          </ModalFooter>
        </>
      </ModalContent>
    </Modal>
  );
};
