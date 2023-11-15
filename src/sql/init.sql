create table if not exists Hardware(
    HardwareID int(10) primary key auto_increment,
    HardwareName varchar(255) not null,
    Category varchar(50),
    Description text,
    Status enum('保留','正常','占用','非正常'),
    Location varchar(100)
);
create table if not exists Software(
    SoftwareID int(10) primary key auto_increment,
    SoftwareName varchar(255) not null,
    Version varchar(50),
    Description text,
    Status enum('正常','非正常'),
    Location varchar(100)
);
create table if not exists Lab(
    LabID int(10) primary key auto_increment,
    LabName varchar(255) not null,
    Description text,
    Status enum('正常','停用','占用')
);
create table if not exists HardwareMaintenance(
    MaintenanceProcessID int(10) primary key auto_increment,
    HardwareID int(10),
    IssueDescription text,
    SolutionDescription text,
    MaintenanceDate date,
    Cost decimal(10,2),
    Status enum('已完成','待处理'),
    foreign key (HardwareID) references Hardware(HardwareID)
);
create table if not exists SoftwareMaintenance(
    MaintenanceProcessID int(10) primary key auto_increment,
    SoftwareID int(10),
    IssueDescription text,
    SolutionDescription text,
    MaintenanceDate date,
    Cost decimal(10,2),
    Status enum('已完成','待处理'),
    foreign key (SoftwareID) references Software(SoftwareID)
);
create table if not exists LabMaintenance(
    MaintenanceProcessID int(10) primary key auto_increment,
    LabID int(10),
    IssueDescription text,
    SolutionDescription text,
    MaintenanceDate date,
    Cost decimal(10,2),
    Status enum('已完成','待处理'),
    foreign key (LabID) references Lab(LabID)
);
create table if not exists User(
    User_id int primary key auto_increment,
    Username varchar(255),
    Pwd varchar(255),
    Userrole varchar(50),
    Full_name varchar(255),
    Email varchar(255),
    Phone_number varchar(20),
    Entry_date date
);
