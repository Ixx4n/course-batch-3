-- USING MS SQL SERVER

-- training.dbo.EXERCISE definition

-- Drop table

-- DROP TABLE training.dbo.EXERCISE GO

CREATE TABLE training.dbo.EXERCISE (
	id int IDENTITY(0,1) NOT NULL,
	title varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	description varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	user_id int NOT NULL,
	CONSTRAINT EXERCISE_PK PRIMARY KEY (id)
) GO;


-- training.dbo.EXERCISE foreign keys

ALTER TABLE training.dbo.EXERCISE ADD CONSTRAINT EXERCISE_FK FOREIGN KEY (user_id) REFERENCES training.dbo.[USER](id) GO;

-- training.dbo.[USER] definition

-- Drop table

-- DROP TABLE training.dbo.[USER] GO

CREATE TABLE training.dbo.[USER] (
	id int IDENTITY(0,1) NOT NULL,
	name varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	email varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	password varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	no_hp varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	created_at datetime NOT NULL,
	updated_at datetime NOT NULL,
	CONSTRAINT USER_PK PRIMARY KEY (id)
) GO;

-- training.dbo.ANSWER definition

-- Drop table

-- DROP TABLE training.dbo.ANSWER GO

CREATE TABLE training.dbo.ANSWER (
	id int IDENTITY(0,1) NOT NULL,
	exercise_id int NOT NULL,
	question_id int NOT NULL,
	user_id int NOT NULL,
	answer varchar(1) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	created_at datetime NOT NULL,
	updated_at datetime NOT NULL,
	CONSTRAINT ANSWER_PK PRIMARY KEY (id)
) GO;


-- training.dbo.ANSWER foreign keys

ALTER TABLE training.dbo.ANSWER ADD CONSTRAINT ANSWER_FK FOREIGN KEY (exercise_id) REFERENCES training.dbo.EXERCISE(id) GO
ALTER TABLE training.dbo.ANSWER ADD CONSTRAINT ANSWER_FK_1 FOREIGN KEY (user_id) REFERENCES training.dbo.[USER](id) GO
ALTER TABLE training.dbo.ANSWER ADD CONSTRAINT ANSWER_FK_2 FOREIGN KEY (question_id) REFERENCES training.dbo.QUESTION(id) GO;

-- training.dbo.QUESTION definition

-- Drop table

-- DROP TABLE training.dbo.QUESTION GO

CREATE TABLE training.dbo.QUESTION (
	id int IDENTITY(0,1) NOT NULL,
	exercise_id int NOT NULL,
	body varchar(MAX) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	option_a varchar(MAX) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	option_b varchar(MAX) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	option_c varchar(MAX) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	option_d varchar(MAX) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	correct_answer varchar(1) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,
	score int NOT NULL,
	user_id int NOT NULL,
	created_at datetime NOT NULL,
	updated_at datetime NOT NULL,
	CONSTRAINT QUESTION_PK PRIMARY KEY (id)
) GO;


-- training.dbo.QUESTION foreign keys

ALTER TABLE training.dbo.QUESTION ADD CONSTRAINT QUESTION_FK FOREIGN KEY (exercise_id) REFERENCES training.dbo.EXERCISE(id) GO
ALTER TABLE training.dbo.QUESTION ADD CONSTRAINT QUESTION_FK_1 FOREIGN KEY (user_id) REFERENCES training.dbo.[USER](id) GO;