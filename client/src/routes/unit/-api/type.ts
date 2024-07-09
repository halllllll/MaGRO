import type { Subunit, Unit } from '@/entity/Unit';
import type { User } from '@/entity/User';

export type Auth = {
  userId: string | undefined;
  idToken: string | undefined;
};

// goで構造体を作ってるが時間がないので手書きする
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
  current_user: User;
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

// goで構造体を定義しているが時間がないので手書きする
export type TargetUsers = {
  user_id: string;
  user_account: string;
};

export type RepassRequest = {
  auth: Auth;
  unitId: number;
  current_user: User;
  target_user: TargetUsers[];
};

export type RepassResultData = {
  user: User;
  status: 'error' | 'success';
  message: string;
  password: string;
};

export type RepassResponse =
  | {
      status: 'error';
      message: string;
    }
  | {
      status: 'success';
      body: RepassResultData[];
    };
