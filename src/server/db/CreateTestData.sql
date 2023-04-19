DROP TABLE TestSchema.Users;
GO

DROP SCHEMA TestSchema;
GO

CREATE SCHEMA TestSchema;
GO

CREATE TABLE TestSchema.Users (
  Id           INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
  Name         NVARCHAR(50),
  Email        NVARCHAR(50),
  Password     NVARCHAR(50)
);
GO

INSERT INTO TestSchema.Users (Name, Email, Password) VALUES
  (N'Joe', N'yourmom@hotmoms.com', N'69420suckmyballs'),
  (N'Bob', N'email@email.com', N'password');
GO

CREATE TABLE TestSchema.Items (
  Id            INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
  Name          NVARCHAR(50),
  Description   NVARCHAR(50),
  UserId        NVARCHAR(50),
  
);
GO

SELECT * FROM TestSchema.Users;
GO