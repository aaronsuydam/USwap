DROP TABLE TestSchema.Users;
GO

DROP TABLE TestSchema.Items;
GO

DROP SCHEMA TestSchema;
GO

CREATE SCHEMA TestSchema;
GO

CREATE TABLE TestSchema.Users (
  Id           INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
  Name         NVARCHAR(50),
  Email        NVARCHAR(50),
  Password     NVARCHAR(60)
);
GO

CREATE TABLE TestSchema.Items (
  Id            INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
  Name          NVARCHAR(50),
  Description   NVARCHAR(50),
  UserId        NVARCHAR(50),
  Image         VARBINARY(MAX)
);
GO