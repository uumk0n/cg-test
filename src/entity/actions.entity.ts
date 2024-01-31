import { Entity, Column, PrimaryGeneratedColumn, ManyToOne, JoinColumn } from 'typeorm';
import { User } from './user.entity';

@Entity()
export class Actions {
  @PrimaryGeneratedColumn()
  id: number;

  @ManyToOne(() => User, (user) => user.actions)
  @JoinColumn({ name: 'user_id' })
  user: User;

  @Column({ name: 'action_time' })
  actionTime: number;

  @Column({ name: 'request_result' })
  requestResult: number;

  @Column({ name: 'temp_c', nullable: true })
  tempC: number;
}
