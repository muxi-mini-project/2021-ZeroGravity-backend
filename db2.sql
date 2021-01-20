DROP DATABASE IF EXISTS `muxi_sharing`;
 
CREATE DATABASE `muxi_sharing`;

USE `muxi_sharing`;

-- 用户信息
CREATE TABLE `user` (
  `id`             INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id`        VARCHAR(15)  NOT NULL COMMENT   "学生学号" ,
  `user_name`      VARCHAR(20)  NOT NULL COMMENT   "用户昵称" ,
  `password`       VARCHAR(15)  NOT NULL COMMENT   "用户密码（一站式用户密码）" ,
  `signture`       VARCHAR(150)  NULL COMMENT   "用户个性签名" ,
  `image_url`      VARCHAR(255)  NULL COMMENT   "头像地址" ,
  `background_url` VARCHAR(255)  NULL COMMENT   "背景地址" ,
  `fans_num`       INT           NULL DEFAULT 0 COMMENT "粉丝数",
  `following_num`  INT           NULL DEFAULT 0 COMMENT "关注的人的数量",
	  
   PRIMARY KEY(`id`) ,
   KEY `user_id`(`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

-- 文件信息
CREATE TABLE `file` (
  `file_id`        INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `file_url`       VARCHAR(255) NOT NULL COMMENT  "文件存储地址" ,
  `file_name`      VARCHAR(30)  NOT NULL COMMENT  "文件标题" ,
  `format`         VARCHAR(10)  NOT NULL COMMENT  "文件格式/(word/txt/pdf/...)" ,
  `content`        VARCHAR(100) NOT NULL COMMENT  "文件内容介绍",
  `subject`        VARCHAR(20)  NOT NULL COMMENT  "学科" ,
  `college`        VARCHAR(20)  NOT NULL COMMENT  "学院" ,
  `type`           VARCHAR(20)  NOT NULL COMMENT  "文件类型/(复习资料/历年真题/...)",
  `grade`          FLOAT        NOT NULL DEFAULT 0 COMMENT "评分" ,
  `like_num`       INT          NOT NULL DEFAULT 0 COMMENT "点赞数" ,
  `collect_num`    INT          NOT NULL DEFAULT 0 COMMENT "收藏数" ,
  `download_num`   INT          NOT NULL DEFAULT 0 COMMENT "下载数" ,
  `scored`         INT          NOT NULL DEFAULT 0 COMMENT "已评论人数" ,

  PRIMARY KEY (`file_id`) ,
  KEY `file_name` (`file_name`) ,
  KEY `format` (`format`) ,
  KEY `type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

-- 关注中间表
CREATE TABLE `following_fans` ( 
  `id`                INT UNSIGNED NOT NULL AUTO_INCREMENT ,
  `following_id`      VARCHAR(15)          NOT NULL COMMENT "主体用户学号" ,
  `fans_id`           VARCHAR(15)           NOT NULL COMMENT "粉丝学号" ,

  PRIMARY KEY (`id`) ,
  KEY `following_id` (`following_id`),
  KEY `fans_id` (`fans_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

-- 上传中间表
CREATE TABLE `file_uploader` (
  `id`             INT UNSIGNED  NOT NULL AUTO_INCREMENT ,
  `uploader_id`  VARCHAR(15)    NOT NULL COMMENT "上传文件的用户ID" ,
  `file_id`      INT    NOT NULL COMMENT "被上传的文件ID" ,
  `upload_time`    varchar(30)      NOT NULL COMMENT "上传时间" ,

  PRIMARY KEY (`id`) ,
  KEY `uploader_id` (`uploader_id`) ,
  KEY `file_id` (`file_id`) ,
  KEY `upload_time` (`upload_time`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

-- 下载中间表
CREATE TABLE `file_downloader` (
  `id`             INT UNSIGNED   NOT NULL AUTO_INCREMENT ,
  `downloader_id`  VARCHAR(15)   NOT NULL COMMENT "下载文件的用户ID" ,
  `file_id`      INT    NOT NULL COMMENT "被下载的文件ID" ,
  `download_time`    varchar(30)     NOT NULL COMMENT "下载时间" ,

  PRIMARY KEY (`id`) ,
  KEY `downloader_id` (`downloader_id`) ,
  KEY `file_id` (`file_id`) ,
  KEY `download_time` (`download_time`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

-- 收藏中间表
CREATE TABLE `file_collecter` (
  `id`              INT UNSIGNED   NOT NULL AUTO_INCREMENT ,
  `collecter_id`  VARCHAR(15)     NOT NULL COMMENT "收藏文件的用户ID" ,
  `file_id`       INT    NOT NULL COMMENT "被收藏的文件ID" ,
  `collect_time`    VARCHAR(30)       NOT NULL COMMENT "收藏时间" ,
  `collectlist_id`  INT  NOT NULL COMMENT "收藏夹id" ,

  PRIMARY KEY (`id`) ,
  KEY `collecter_id` (`collecter_id`) ,
  KEY `file_id` (`file_id`) ,
  KEY `collect_time` (`collect_time`),
  KEY `collectlist_id` (`collectlist_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

-- 留言中间表
CREATE TABLE `message` (
  `id`             INT UNSIGNED  NOT NULL AUTO_INCREMENT ,
  `writer_id`    VARCHAR(15)    NOT NULL COMMENT "写留言的人id" ,
  `host_id`      VARCHAR(15)   NOT NULL COMMENT "留言版主人id" ,
  `write_time`     varchar(30)      NOT NULL COMMENT "写留言时间" ,
  `content`      VARCHAR(256)  NOT NULL COMMENT "留言内容",

  PRIMARY KEY (`id`) ,
  KEY `host_id` (`host_id`) 
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

-- 点赞中间表
CREATE TABLE `likes` (
  `id`             INT UNSIGNED  NOT NULL AUTO_INCREMENT ,
  `user_id`    VARCHAR(15)    NOT NULL COMMENT "点赞的人id" ,
  `file_id`      INT   NOT NULL COMMENT "文件id" ,

  PRIMARY KEY (`id`) ,
  KEY `user_id` (`user_id`) 
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `score` (
  `id`             INT UNSIGNED  NOT NULL AUTO_INCREMENT ,
  `user_id`        VARCHAR(15)   NOT NULL COMMENT "点赞的人id" ,
  `file_id`        INT   NOT NULL COMMENT "文件id" ,
  `score`          INT   NOT NULL COMMENT  "文件单次评分" ,

  PRIMARY KEY (`id`) ,
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `collect_list` (
  `collectlist_id`     INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "收藏夹编号" ,
  `collectlist_name`   VARCHAR(30)   NOT NULL COMMENT  "收藏夹昵称" ,
  `user_id`       VARCHAR(15) NOT NULL COMMENT "收藏夹所属用户" ,

  PRIMARY KEY (`collectlist_id`) ,
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;
