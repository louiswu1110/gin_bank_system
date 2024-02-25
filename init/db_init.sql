CREATE DATABASE meepshop_bank;

Use meepshop_bank;
create table member
(
    id  bigint   not null comment '會員ID' primary key,
    nickname varchar(100) not null comment '暱稱',
    username varchar(50) not null comment '姓名',
    added_time  datetime(6) default CURRENT_TIMESTAMP(6) not null comment '創建時間',
    UNIQUE KEY `idx_username` (`username`)
) comment '會員';

create table `transaction`
(
    id bigint not null comment '交易序號' primary key,
    member_id               bigint                                   not null comment '會員ID',
    source_type             tinyint                                  not null comment '來源類型 0 會員提款, 1 會員存款, 2 轉帳出款, 3 轉帳入款, 4 人工入款',
    amount                  decimal(18, 2)                           not null comment '金額',
    current_balance         decimal(18, 2)                           not null comment '變動前金額',
    changed_balance         decimal(18, 2)                           not null comment '變動後金額',
    added_time              datetime(6) default CURRENT_TIMESTAMP(6) not null comment '創建時間',
    remarks                 tinytext                                 null comment '備註'
) comment '交易記錄';

create table member_account
(
    member_id     bigint                                      not null comment '會員id'
        primary key,
    balance       decimal(18, 2) default 0.00                 not null comment '餘額',
    added_time    datetime(6)    default CURRENT_TIMESTAMP(6) not null comment '創建時間',
    updated_time  datetime(6)                                 null on update CURRENT_TIMESTAMP(6) comment '修改時間'
) comment '會員帳戶';


INSERT INTO `member` (`id`, `nickname`, `username`, `added_time`)
VALUES ((SELECT CONCAT(FLOOR(UNIX_TIMESTAMP() * 1000), FLOOR(10000 + RAND() * 90000))),
        '小明明', 'ming', NOW());

INSERT INTO `member` (`id`, `nickname`, `username`, `added_time`)
VALUES ((SELECT CONCAT(FLOOR(UNIX_TIMESTAMP() * 1000), FLOOR(10000 + RAND() * 90000))),
        '大熊熊', 'bear', NOW());

INSERT INTO `member_account` (`member_id`, `balance`, `added_time`, `updated_time`)
VALUES ((SELECT `id` FROM `member` WHERE username = 'bear'), 0, NOW(), NOW());

INSERT INTO `member_account` (`member_id`, `balance`, `added_time`, `updated_time`)
VALUES ((SELECT `id` FROM `member` WHERE username = 'ming') , 0, NOW(), NOW());


