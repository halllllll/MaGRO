import type { Subunit, Unit } from '@/entity/Unit';
import type { User } from '@/entity/User';

export type Auth = {
  userId: string | undefined;
  idToken: string | undefined;
};

// goで構造体を作っておるが時間がないので手書きする
// type Result = 'success' | 'error';

type UsersWithSubgroups = {
  user: User;
  subunit_ids: number[];
};

type OperatorSubunit = {
  operators: string[];
  subunit: Subunit;
};

export type SuccessData = {
  unit: Unit;
  current_usr: User;
  user_count: number;
  user_groups: UsersWithSubgroups[];
  subunit_count: number;
  subunit_groups: OperatorSubunit[];
};

export type UsersSubunitResponse =
  | {
      status: 'error';
      message: string;
    }
  | {
      status: 'success';
      data: SuccessData;
    };
