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
} from '@chakra-ui/react';
import type { FC } from 'react';
import { type SubmitHandler, useFormContext } from 'react-hook-form';
import type { SchemaType } from './schema';
import type { User } from '@/entity/User';

type Props = {
  isOpen: boolean;
  onClose: () => void;
  onConfirm: SubmitHandler<SchemaType>;
  data: Map<string, User>;
};

export const ConfirmModal: FC<Props> = (props) => {
  const { isOpen, onClose, data, onConfirm } = props;

  // checkbox formで選択したユーザーIDからユーザー情報を構築
  const methods = useFormContext<SchemaType>();
  const userIds = methods.getValues('user_ids') ?? [];

  // 送信時はsubmitとモーダル閉じるのをやりたい
  // 送信操作はFormで行う

  return (
    <Modal
      isOpen={isOpen}
      size={'xl'}
      onClose={onClose}
      closeOnOverlayClick={false}
      allowPinchZoom={true}
      scrollBehavior={'inside'}
    >
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>MaGRO Confirm</ModalHeader>
        <>
          <ModalCloseButton />
          <ModalBody pb={6}>
            <Box>
              <Text mb={5}>次のアカウントのパスワードをリセットします</Text>
              <Alert status="error">
                <AlertIcon />
                <AlertTitle>{''}</AlertTitle>
                <Box>
                  <Text as={'b'}>ATTENTION</Text>
                  <UnorderedList styleType="'- '">
                    <ListItem>
                      <AlertDescription>
                        パスワードリセット後、ユーザーは現在のパスワードを使用してログインすることができません。
                      </AlertDescription>
                    </ListItem>
                    <ListItem>
                      <AlertDescription>
                        新しいパスワードをユーザーに安全に通知する方法を確認してください。
                      </AlertDescription>
                    </ListItem>
                    <ListItem>
                      <AlertDescription>
                        新しいパスワードはサーバーで保存されません。
                      </AlertDescription>
                    </ListItem>
                    <AlertDescription>
                      <ListItem>この操作は取り消しできません。</ListItem>
                    </AlertDescription>
                  </UnorderedList>
                </Box>
              </Alert>
            </Box>
            <Box>
              <TableContainer w={'100%'} whiteSpace={'unset'} overflowX="unset" overflowY="unset">
                <Table variant="simple">
                  <Tbody>
                    {userIds
                      .filter((v) => !!v) // false避け
                      .filter((v) => v !== undefined && v !== null) // なんか後続の処理でなぜかvが string | undefined | nullなため
                      .filter((vv, ii, ss) => ss.indexOf(vv) === ii)
                      .map((v) => {
                        const user = data.get(v);
                        return (
                          <Tr key={user?.user_id}>
                            <Td>{user?.user_displayname}</Td>
                            <Td>{user?.user_sortkey}</Td>
                          </Tr>
                        );
                      })}
                  </Tbody>
                </Table>
              </TableContainer>
            </Box>
            <VStack spacing={4} />
          </ModalBody>
          <ModalFooter>
            <form onSubmit={methods.handleSubmit(onConfirm)}>
              <Button
                type="submit"
                colorScheme="red"
                mr={3}
                loadingText="submitting..."
                spinnerPlacement="start"
                onClick={onClose}
              >
                OK
              </Button>
              <Button onClick={onClose}>Cancel</Button>
            </form>
          </ModalFooter>
        </>
      </ModalContent>
    </Modal>
  );
};
