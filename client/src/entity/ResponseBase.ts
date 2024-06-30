export type RespBase =
  | {
      status: 'error';
      message: string;
    }
  | {
      status: 'success';
    };
