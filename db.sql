drop database if exists `ZeroGravity`;

create database `ZeroGravity`;

use  `ZeroGravity`;


-- 用户信息表(user)
create table `tbl_user`(
   `id`               int           not null  AUTO_INCREMENT             comment "用户id" ,       
   `account_password` varchar(20)       null                             comment "账户密码",
   `account`          varchar(20)       null                             comment "Q Q账号",
   `nickname`         varchar(20)       null                             comment "昵称",
   `avatar`           varchar(255)      null                             comment "头像",
   `energy`           int               null  default 0                  comment "能量值"  ,   

-- 添加约束
primary key                         (`id`),
key  `account_password`             (`account_password`),
key  `account`                      (`account`)
FULLTEXT (`nickname`) WITH PARSER ngram
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 想法表(idea)
create table `tbl_idea` (
    `idea_id`         int           not null AUTO_INCREMENT comment "想法id",
    `content`         varchar(255)      null comment "想法内容",
    `releaseDate`     date              null comment "发布日期",
    `publisher_id`    varchar(255)      null comment "发布者id",
    `likes_sum`       int               null comment "点赞数",
    `comments_sum`    int               null comment "评论数",
    `privacy`         tinyint(1)      DEFAULT NULL,
-- 添加约束
primary key                         (`idea_id`)
FULLTEXT (`content`) WITH PARSER ngram
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 评论表(comments)
create table `tbl_comments`(
    `id`               int          not null AUTO_INCREMENT comment "评论id" ,
    `commenter_id`     int              null comment "评论者id",
    `commented_id`     int              null comment "idea_id",
    `likes_sum`        int              null comment "赞数",
    `release_date`     date             null comment "发布日期",
    `content`          varchar(255)     null comment "内容",
-- 添加约束
primary key                          (`id`),
key           `commenter_id`         (`commenter_id`),
key           `commented_id`         (`commented_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 收藏表(favorite_records)
create table `tbl_favorite_records`(
    `id`               int           not null AUTO_INCREMENT comment "收藏记录序号",
    `collector_id`     int               null comment "收藏者id",
    `idea_id`          int               null comment "想法id",
-- 添加约束
primary key                           (`id`),
key            `collector_id`         (`collector_id`),
key            `idea_id`              (`idea_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 想法点赞记录表(like_record_idea)
create table `tbl_like_record_idea`(
    `id`               int            not null AUTO_INCREMENT comment "点赞记录序号",
    `idea_id`          int                     comment "想法id",
    `likers_id`        int                     comment "点赞者id",
    `beliked_id`       int                     comment "被点赞者id",
-- 添加约束
primary key                           (`id`),
key             `idea_id`             (`idea_id`),
key             `likers_id`           (`likers_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 评论点赞记录表(comment_record_idea)
create table `tbl_like_record_comment`(
    `id`               int            not null AUTO_INCREMENT comment "点赞记录序号",
    `comment_id`       int                null comment "评论id",
    `likers_id`        int                null comment "点赞者id",
    `beliked_id`       int                null comment "被点赞者id",
-- 添加约束
primary key                            (`id`),
key                `comment_id`        (`comment_id`),
key                `likers_id`         (`likers_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

DROP TABLE IF EXISTS histories;
create table histories(
    id      int          not null auto_increment ,
    name    varchar(100) not null ,
    user_id varchar(100) not null ,
    primary key (id)
)ENGINE=InnoDB;
