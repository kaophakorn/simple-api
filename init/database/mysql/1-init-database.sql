SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Create Database
-- ----------------------------
CREATE DATABASE IF NOT EXISTS `main` DEFAULT CHARACTER SET utf8mb4;

-- ----------------------------
-- Use Database
-- ----------------------------
USE `main`;

-- ----------------------------
-- Table structure for league
-- ----------------------------
CREATE TABLE `league`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of league
-- ----------------------------
BEGIN;
INSERT INTO `league` (`id`, `name`, `created_at`, `updated_at`) VALUES (1, 'Premier League', '2022-02-10 14:16:42', '2022-02-10 14:16:42');
COMMIT;

-- ----------------------------
-- Table structure for player
-- ----------------------------
CREATE TABLE `player`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `team_id` int(10) UNSIGNED NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `middlename` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `lastname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `pos` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of player
-- ----------------------------
BEGIN;
INSERT INTO `player` (`id`, `team_id`, `name`, `middlename`, `lastname`, `pos`, `created_at`, `updated_at`) VALUES (1, 1, 'Salah', '', 'Mohamed', 'RW', '2022-02-10 13:53:03', '2022-02-10 13:53:03');
COMMIT;

-- ----------------------------
-- Table structure for team
-- ----------------------------
CREATE TABLE `team`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `league_id` int(10) UNSIGNED NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of team
-- ----------------------------
BEGIN;
INSERT INTO `team` (`id`, `league_id`, `name`, `created_at`, `updated_at`) VALUES (1, 1, 'Liverpool', '2022-02-10 14:10:26', '2022-02-10 14:10:26');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
