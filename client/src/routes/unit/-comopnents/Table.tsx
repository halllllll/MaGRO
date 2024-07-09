import type { FC } from 'react';
import type { SuccessData } from '../-api/type';
import {
  Button,
  Checkbox,
  Text,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
  Flex,
} from '@chakra-ui/react';
import type { Subunit } from '@/entity/Unit';
import type { User } from '@/entity/User';
import { useFormContext } from 'react-hook-form';
import type { SchemaType } from './schema';

type Props = {
  data: SuccessData;
  isSending?: boolean;
};

export const UsersSubunitsList: FC<Props> = ({ data, isSending }) => {
  // subunitごとに一覧するために使いやすそうなデータ構造にする
  // MapはキーがObjectでも厳密等価で比較できるっぽい
  const SubunitIdMap = new Map<number, Subunit>();
  const SubunitUsersMap = new Map<Subunit, User[]>();
  // ひとまず空にして、あとから入れ直す
  // sort -> nameでソートしながらやる
  const sorted_subunit_groups = data.subunit_groups.toSorted((a, b) => {
    return a.subunit.name < b.subunit.name ? -1 : 1;
  });
  for (const su of sorted_subunit_groups) {
    SubunitIdMap.set(su.subunit.subunit_id, su.subunit);
    SubunitUsersMap.set(su.subunit, []);
  }

  for (const u of data.user_groups) {
    for (const sid of u.subunit_ids) {
      const targetSubunit = SubunitIdMap.get(sid);
      if (!targetSubunit) continue; // いったん無視
      const tmpUsers: User[] = SubunitUsersMap.get(targetSubunit) ?? [];
      const updateUsers = [...tmpUsers, u.user];
      SubunitUsersMap.set(targetSubunit, updateUsers);
    }
  }

  // なんかvalidationがきかない。いったんsubunitなしにユーザーデータだけのやつでやってみる
  // 平だったらできたので、たぶん二重でmapでやってるのがだめ？
  // ということでテーブルの行用のデータを作り、一層で回す
  // type RowType = 'user' | 'subunit';
  type Row =
    | {
        rowType: 'user';
        data: User;
      }
    | {
        rowType: 'subunit';
        data: {
          subunit: Subunit;
          memberCount: number;
        };
      };
  const row: Row[] = [];

  for (const sg of sorted_subunit_groups) {
    const subunitId = sg.subunit.subunit_id;
    const targetSubunit = SubunitIdMap.get(subunitId);
    if (!targetSubunit) continue;
    const users = SubunitUsersMap.get(targetSubunit);
    if (!users) continue;
    row.push({ rowType: 'subunit', data: { subunit: targetSubunit, memberCount: users.length } });
    for (const u of users) {
      row.push({ rowType: 'user', data: u });
    }
  }

  // rfh
  const { register, formState, watch } = useFormContext<SchemaType>();
  const w = watch('user_ids');
  const validCount = w?.filter((v) => !!v).filter((e, i, s) => s.indexOf(e) === i).length ?? 0;
  return (
    <>
      {/* <Text>{JSON.stringify(w)}</Text>  デバッグ用　*/}
      <Flex
        justifyContent={'right'}
        align={'center'}
        minH={4}
        backgroundColor={'Menu'}
        borderWidth={'2px'}
        border={'Background'}
        mb={4}
        mr={10}
        gap={4}
      >
        {`${validCount} 件選択中`}
        <Button
          type={'submit'}
          isDisabled={!validCount || !formState.isValid || formState.isSubmitting || isSending}
          isLoading={isSending}
          variant={'solid'}
          colorScheme={'teal'}
        >
          confirm
        </Button>
      </Flex>
      <TableContainer w={'100%'} whiteSpace={'unset'} overflowX="unset" overflowY="unset">
        <Table size={'sm'}>
          <Thead position="sticky" zIndex="docked">
            <Tr>
              <Th>{''}</Th>
              <Th>id</Th>
              <Th>name</Th>
              <Th>kana</Th>
            </Tr>
          </Thead>
          <Tbody>
            {/* {sorted_subunit_groups.map((v) => {
              const subunitId = v.subunit.subunit_id;
              const targetSubunit = SubunitIdMap.get(subunitId);
              if (!targetSubunit) return null;
              const users = SubunitUsersMap.get(targetSubunit);
              if (!users) return null;
              return (
                <>
                  <Tr key={v.subunit.subunit_id}>
                    <Td fontSize={'xl'} colSpan={9999} height={'3.5rem'}>
                      {targetSubunit.name}
                    </Td>
                  </Tr>

                  {users.map((u, idx) => {
                    return (
                      <Tr key={`${idx}_${u.user_id}`} _hover={{ bg: 'gray.100' }} height={'3rem'}>
                        <Td>
                          <Checkbox
                            {...register(`users.${idx}`)}
                            name={`users.${idx}`}
                            value={u.user_id}
                          />
                        </Td>
                        <Td>{u.user_name}</Td>
                        <Td>
                          <Box>{u.user_displayname}</Box>
                        </Td>
                        <Td>
                          <Box>{u.user_sortkey}</Box>
                        </Td>
                      </Tr>
                    );
                  })}
                </>
              );
            })} */}
            {/** ↓これはできた　 */}
            {/* {data.user_groups.map((v, idx) => {
              return (
                <Tr key={`${idx}_${v.user.user_id}`}>
                  <Td>
                    <Checkbox
                      {...register(`users.${idx}`)}
                      name={`users.${idx}`}
                      value={v.user.user_id}
                    />
                  </Td>
                  <Td>{v.user.user_name}</Td>
                  <Td>{v.user.user_displayname}</Td>
                </Tr>
              );
            })} */}
            {row.map((r, idx) => {
              return r.rowType === 'subunit' ? (
                <Tr key={`${idx}_${r.data.subunit.name}`}>
                  <Td fontSize={'xl'} colSpan={9999} height={'3.5rem'}>
                    <Flex gap={4}>
                      <Text>{r.data.subunit.name}</Text>
                      <Text fontSize={'md'}>{`(count ${r.data.memberCount})`}</Text>
                    </Flex>
                  </Td>
                </Tr>
              ) : (
                <Tr key={`${idx}_${r.data.user_id}`} _hover={{ bg: 'gray.100' }} height={'3rem'}>
                  <Td>
                    <Checkbox
                      {...register(`user_ids.${idx}`)}
                      name={`user_ids.${idx}`}
                      value={r.data.user_id}
                    />
                  </Td>
                  <Td>{r.data.user_name}</Td>
                  <Td>{r.data.user_displayname}</Td>
                </Tr>
              );
            })}
          </Tbody>
        </Table>
      </TableContainer>
    </>
  );
};
