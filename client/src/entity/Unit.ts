export type Unit = {
  unit_id: number;
  name: string;
};

export type Subunit = {
  subunit_id: number;
  name: string;
  isPublic: boolean;
  created: Date;
  modifed: Date;
};
