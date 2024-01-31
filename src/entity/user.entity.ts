import { Entity, Column, PrimaryGeneratedColumn, OneToMany } from 'typeorm';
import { Actions } from './actions.entity';

@Entity()
export class User {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ unique: true })
  login: string;

  @Column()
  hashedPassword: string;

  @Column()
  fio: string;

  @Column()
  apiToken: string;

  @OneToMany(() => Actions, (actions) => actions.user)
  actions: Actions[];
}
