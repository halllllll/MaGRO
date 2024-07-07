import * as yup from 'yup';

// よくわからないが今のコンポーネント構造だとfalseではなくnullになるやつがある
export const schema = yup.object().shape({
  user_ids: yup.array().of(yup.string().nullable()).min(1).required(),

  // users: yup
  //   .array()
  //   .of(
  //     yup.object().shape({
  //       id: yup.string().required(),
  //     }),
  //   )
  //   .min(1)
  //   .required(),

  // なんか入れ子でmapでreturnするコンポーネント構造だとスキーマもそうなるっぽい？
  // users: yup.array().of(yup.array().of(yup.string().required())).min(1).required(),
});

export type SchemaType = yup.InferType<typeof schema>;
