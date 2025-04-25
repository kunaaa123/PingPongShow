-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 25, 2025 at 08:05 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.0.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `pingshow`
--

-- --------------------------------------------------------

--
-- Table structure for table `all_matches`
--

CREATE TABLE `all_matches` (
  `id` int(11) NOT NULL,
  `event_time` datetime NOT NULL,
  `event_type` varchar(50) NOT NULL,
  `player` varchar(50) NOT NULL,
  `power` int(11) NOT NULL,
  `goroutine` varchar(100) NOT NULL,
  `match_number` int(11) NOT NULL,
  `turn` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `all_matches`
--

INSERT INTO `all_matches` (`id`, `event_time`, `event_type`, `player`, `power`, `goroutine`, `match_number`, `turn`, `created_at`) VALUES
(1, '2025-04-25 05:39:23', 'start_game', 'System', 0, 'main', 11, 0, '2025-04-25 05:42:27'),
(2, '2025-04-25 05:39:23', 'wake_up', 'B', 0, 'B-1745559563317203800', 11, 2, '2025-04-25 05:42:27'),
(3, '2025-04-25 05:39:23', 'wake_up', 'A', 0, 'A-1745559563317203800', 11, 1, '2025-04-25 05:42:27'),
(4, '2025-04-25 05:39:23', 'ping', 'A', 61, 'A-1745559563317203800', 11, 1, '2025-04-25 05:42:27'),
(5, '2025-04-25 05:39:23', 'table_response', 'Table', 54, 'Table', 11, 0, '2025-04-25 05:42:27'),
(6, '2025-04-25 05:39:23', 'wake_up', 'B', 54, 'B-1745559563317203800', 11, 2, '2025-04-25 05:42:27'),
(7, '2025-04-25 05:39:23', 'pong', 'B', 7, 'B-1745559563317203800', 11, 2, '2025-04-25 05:42:27'),
(8, '2025-04-25 05:39:23', 'lose', 'B', 7, 'B-1745559563317203800', 11, 2, '2025-04-25 05:42:27'),
(9, '2025-04-25 05:39:23', 'game_over', 'System', 0, 'main', 11, 2, '2025-04-25 05:42:27'),
(10, '2025-04-25 05:39:25', 'start_game', 'System', 0, 'main', 12, 0, '2025-04-25 05:42:27'),
(11, '2025-04-25 05:39:25', 'wake_up', 'A', 0, 'A-1745559565519380600', 12, 1, '2025-04-25 05:42:27'),
(12, '2025-04-25 05:39:25', 'wake_up', 'B', 0, 'B-1745559565519380600', 12, 2, '2025-04-25 05:42:27'),
(13, '2025-04-25 05:39:25', 'ping', 'A', 80, 'A-1745559565519380600', 12, 1, '2025-04-25 05:42:27'),
(14, '2025-04-25 05:39:25', 'table_response', 'Table', 71, 'Table', 12, 0, '2025-04-25 05:42:27'),
(15, '2025-04-25 05:39:25', 'wake_up', 'B', 71, 'B-1745559565519380600', 12, 2, '2025-04-25 05:42:27'),
(16, '2025-04-25 05:39:25', 'pong', 'B', 78, 'B-1745559565519380600', 12, 2, '2025-04-25 05:42:27'),
(17, '2025-04-25 05:39:25', 'wake_up', 'A', 78, 'A-1745559565519380600', 12, 3, '2025-04-25 05:42:27'),
(18, '2025-04-25 05:39:25', 'ping', 'A', 87, 'A-1745559565519380600', 12, 3, '2025-04-25 05:42:27'),
(19, '2025-04-25 05:39:25', 'table_response', 'Table', 61, 'Table', 12, 0, '2025-04-25 05:42:27'),
(20, '2025-04-25 05:39:25', 'wake_up', 'B', 61, 'B-1745559565519380600', 12, 4, '2025-04-25 05:42:27'),
(21, '2025-04-25 05:39:25', 'pong', 'B', 84, 'B-1745559565519380600', 12, 4, '2025-04-25 05:42:27'),
(22, '2025-04-25 05:39:25', 'wake_up', 'A', 84, 'A-1745559565519380600', 12, 5, '2025-04-25 05:42:27'),
(23, '2025-04-25 05:39:25', 'ping', 'A', 76, 'A-1745559565519380600', 12, 5, '2025-04-25 05:42:27'),
(24, '2025-04-25 05:39:25', 'table_response', 'Table', 55, 'Table', 12, 0, '2025-04-25 05:42:27'),
(25, '2025-04-25 05:39:26', 'wake_up', 'B', 55, 'B-1745559565519380600', 12, 6, '2025-04-25 05:42:27'),
(26, '2025-04-25 05:39:26', 'pong', 'B', 5, 'B-1745559565519380600', 12, 6, '2025-04-25 05:42:27'),
(27, '2025-04-25 05:39:26', 'lose', 'B', 5, 'B-1745559565519380600', 12, 6, '2025-04-25 05:42:27'),
(28, '2025-04-25 05:39:26', 'game_over', 'System', 0, 'main', 12, 6, '2025-04-25 05:42:27'),
(29, '2025-04-25 05:39:28', 'start_game', 'System', 0, 'main', 13, 0, '2025-04-25 05:42:27'),
(30, '2025-04-25 05:39:28', 'wake_up', 'A', 0, 'A-1745559568122895800', 13, 1, '2025-04-25 05:42:27'),
(31, '2025-04-25 05:39:28', 'wake_up', 'B', 0, 'B-1745559568122895800', 13, 2, '2025-04-25 05:42:27'),
(32, '2025-04-25 05:39:28', 'ping', 'A', 93, 'A-1745559568122895800', 13, 1, '2025-04-25 05:42:27'),
(33, '2025-04-25 05:39:28', 'table_response', 'Table', 82, 'Table', 13, 0, '2025-04-25 05:42:27'),
(34, '2025-04-25 05:39:28', 'wake_up', 'B', 82, 'B-1745559568122895800', 13, 2, '2025-04-25 05:42:27'),
(35, '2025-04-25 05:39:28', 'pong', 'B', 49, 'B-1745559568122895800', 13, 2, '2025-04-25 05:42:27'),
(36, '2025-04-25 05:39:28', 'lose', 'B', 49, 'B-1745559568122895800', 13, 2, '2025-04-25 05:42:27'),
(37, '2025-04-25 05:39:28', 'game_over', 'System', 0, 'main', 13, 2, '2025-04-25 05:42:27');

-- --------------------------------------------------------

--
-- Table structure for table `last_played_match`
--

CREATE TABLE `last_played_match` (
  `id` int(11) NOT NULL,
  `match_number` int(11) NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `last_played_match`
--

INSERT INTO `last_played_match` (`id`, `match_number`, `updated_at`) VALUES
(1, 13, '2025-04-25 05:42:27');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `all_matches`
--
ALTER TABLE `all_matches`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `last_played_match`
--
ALTER TABLE `last_played_match`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `all_matches`
--
ALTER TABLE `all_matches`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=38;

--
-- AUTO_INCREMENT for table `last_played_match`
--
ALTER TABLE `last_played_match`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
