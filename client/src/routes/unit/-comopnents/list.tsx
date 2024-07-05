import type { FC } from 'react';
import type { SuccessData } from '../-api/type';
import { Box, Checkbox, Table, TableContainer, Tbody, Td, Th, Thead, Tr } from '@chakra-ui/react';
import type { Subunit } from '@/entity/Unit';
import type { User } from '@/entity/User';

type Props = {
  data: SuccessData;
};

export const UsersSubunitsListX: FC<Props> = ({ data }) => {
  // subunitごとに一覧する
  const SubunitIdMap = new Map<number, Subunit>();
  const SubunitUsersMap = new Map<Subunit, User[]>();
  // ひとまず空にして、あとから入れ直す
  for (const su of data.subunit_groups) {
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

  return (
    <TableContainer w={'100%'} whiteSpace={'unset'} overflowX="unset" overflowY="unset">
      <Table size={'sm'}>
        <Thead position="sticky" zIndex="docked">
          <Tr>
            <Th w="10%">{''}</Th>
            <Th>id</Th>
            <Th>name</Th>
            <Th>kana</Th>
          </Tr>
        </Thead>
        <Tbody>
          {data.subunit_groups.map((v) => {
            const subunitId = v.subunit.subunit_id;
            const targetSubunit = SubunitIdMap.get(subunitId);
            if (!targetSubunit) return <></>;
            const users = SubunitUsersMap.get(targetSubunit);
            if (!users) return <></>;
            return (
              <>
                <Tr key={v.subunit.subunit_id}>
                  {/** だめなのはわかっとる */}
                  <Td fontSize={'xl'} colSpan={9999} height={'8vh'}>
                    {targetSubunit.name}
                  </Td>
                </Tr>
                {users.map((u) => {
                  return (
                    <Tr key={u.user_id} _hover={{ bg: 'gray.100' }} height={'6vh'}>
                      <Td>
                        {' '}
                        <Checkbox size="lg" />
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
          })}
        </Tbody>
      </Table>
    </TableContainer>
  );
};
