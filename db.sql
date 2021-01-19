--此脚本是为“ZeroGravity”所设计的建表文件
--作者：龚娜
--建表时间：2020-01-19 15：30

--用户（user）表
create table tbl_userInformation(
    id int  ,
    accountPassword varchar2(255) not null,
    account varchar2(255) not null,
    nickname varchar2(255) not null,
    avatar varchar2(255),
    energy int 
 --添加约束
primary key (id),
unique (name )

);


--想法表(idea)
create table tbl_userIdea (
    id int ,
    content varchar not null ,
    releaseDate date default sysdate,
    publisherID int ,
    heat int ,
    numberOfComments int ,
--添加约束
primary key(id),
foreign key(publisherID)
);



--评论表(comments)
create table tbl_usercomments(
    id int ,
    commentatorID int ,
    commentOnID int ,
    likes int ,
    releaseDate date default sysdate ,
    content varchar2(255) not null,
--添加约束
primary key (id),
foreign key (commentatorID),
foreign key (commentOnID),
);



--收藏记录表(favoriteRecords)
create table tbl_favoriteRecords(
    id int,
    user_id int,
    idea_id int,
--添加约束
primary key (id),
foreign key (user_id),
foreign key (idea_id),

);



--想法点赞记录表(likeRecordOfIdea)
create table tbl_likeRecordOfIdea (
    id int,
    idea_Id int ,
    likesId int ,
    beLikedPersonID int ,
--添加约束
primary key (id),
foreign key (idea_Id),
foreign key (likesId),
foreign key (beLikedPersonID)

);

--评论点赞记录表(likeRecordOfComments)
create table tbl_likeRecordOfComments(
    id int ,
    commentID int,
    likesId int,
    beLikedPersonID int ,
--添加约束
primary key(id),
foreign key (commentID),
foreign key (likesId),
foreign key (beLikedPersonID)
);
