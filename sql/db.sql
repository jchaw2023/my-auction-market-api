-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        11.5.2-MariaDB - mariadb.org binary distribution
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  12.12.0.7122
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- 导出 auction_market_db 的数据库结构
CREATE DATABASE IF NOT EXISTS `auction_market_db` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;
USE `auction_market_db`;

-- 导出  表 auction_market_db.auctions 结构
CREATE TABLE IF NOT EXISTS `auctions` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `auction_id` varchar(50) DEFAULT NULL COMMENT '拍卖ID(使用snowflake生成ID，避免自增ID)',
  `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '创建者ID',
  `owner_address` varchar(42) DEFAULT NULL COMMENT '当前拥有者地址',
  `contract_auction_id` bigint(20) unsigned NOT NULL COMMENT '合约里面拍卖列表索引',
  `nft_id` varchar(64) DEFAULT NULL COMMENT 'NFT唯一标识',
  `nft_address` varchar(42) NOT NULL COMMENT 'NFT合约地址',
  `token_id` bigint(20) unsigned NOT NULL COMMENT 'NFT的Token ID',
  `contract_name` varchar(255) DEFAULT NULL COMMENT '合约名称',
  `contract_symbol` varchar(64) DEFAULT NULL COMMENT '合约符号',
  `token_uri` text DEFAULT NULL COMMENT 'Token URI',
  `nft_name` varchar(255) DEFAULT NULL COMMENT 'NFT名称',
  `image` text DEFAULT NULL COMMENT 'NFT图片URL',
  `description` text DEFAULT NULL COMMENT 'NFT描述',
  `metadata` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '完整元数据JSON',
  `status` varchar(20) NOT NULL DEFAULT 'pending' COMMENT '状态(pending,active,ended,cancelled)',
  `online_lock` varchar(100) DEFAULT NULL COMMENT 'NFT在线标志 nft_id:1,也作为一个锁字段，解锁就改成其他值',
  `online` bigint(20) DEFAULT NULL COMMENT '1表示在线 其他值表示下线',
  `start_time` datetime NOT NULL COMMENT '开始时间',
  `start_timestamp` bigint(20) NOT NULL DEFAULT 0 COMMENT '开始时间时间戳',
  `end_time` datetime NOT NULL COMMENT '结束时间',
  `end_timestamp` bigint(20) NOT NULL DEFAULT 0 COMMENT '结束时间时间戳',
  `start_price` decimal(65,30) DEFAULT NULL COMMENT '起拍价(单位由PaymentToken指定:0x0=ETH,其他=ERC20代币)',
  `payment_token` varchar(42) DEFAULT NULL COMMENT '起拍价链上交易代币地址(0x0表示ETH,其他地址表示ERC20代币)',
  `start_price_usd` decimal(65,30) DEFAULT NULL COMMENT '起拍价USD',
  `start_price_unit_usd` bigint(20) DEFAULT NULL COMMENT '起拍价USD预言机价格（小数点起拍价USD*10**8）',
  `highest_bidder` varchar(42) DEFAULT NULL COMMENT '最高出价者地址',
  `highest_bid_payment_token` varchar(42) DEFAULT NULL COMMENT '最高出价使用链上交易代币地址(0x0=ETH,其他=ERC20代币,可能与拍卖PaymentToken不同)',
  `highest_bid` decimal(65,30) unsigned DEFAULT NULL COMMENT '最高出价金额(单位由HighestBidPaymentToken指定,可能与拍卖PaymentToken不同)',
  `highest_bid_usd` decimal(65,30) DEFAULT NULL COMMENT '最高出价USD(用于比较不同代币的出价)',
  `highest_bid_unit_usd` bigint(20) DEFAULT NULL COMMENT '最高出价USD预言机价格（小数点最高价USD*10**8）',
  `bid_count` bigint(20) unsigned DEFAULT 0 COMMENT '出价次数',
  `created_at` datetime NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `nft_online_id` (`online_lock`) USING BTREE,
  UNIQUE KEY `auction_id` (`auction_id`),
  KEY `idx_auctions_user_id` (`user_id`),
  KEY `idx_auctions_nft_id` (`nft_id`),
  KEY `idx_auctions_owner` (`owner_address`),
  KEY `idx_auctions_status` (`status`),
  KEY `idx_auctions_nft_address` (`nft_address`),
  KEY `idx_auctions_token_id` (`token_id`),
  KEY `idx_auctions_contract_id` (`contract_auction_id`) USING BTREE,
  KEY `online` (`online`),
  KEY `start_timestamp` (`start_timestamp`),
  KEY `end_timestamp` (`end_timestamp`),
  CONSTRAINT `fk_auctions_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='拍卖表';

-- 数据导出被取消选择。

-- 导出  表 auction_market_db.bids 结构
CREATE TABLE IF NOT EXISTS `bids` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '出价ID',
  `auction_id` varchar(50) DEFAULT NULL COMMENT '拍卖ID',
  `contract_auction_id` bigint(20) unsigned NOT NULL COMMENT '拍卖合约里面的拍卖ID',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '出价者ID',
  `wallet_address` varchar(50) DEFAULT NULL COMMENT '出价者钱包地址',
  `winner` tinyint(1) DEFAULT 0 COMMENT '竞拍获胜者',
  `amount` decimal(20,8) NOT NULL DEFAULT 0.00000000 COMMENT '出价金额(ETH,USDC)',
  `amount_unit` bigint(20) NOT NULL DEFAULT 0 COMMENT '出价金额(wei、usdc最小单位等)',
  `amount_usd` decimal(20,8) DEFAULT NULL COMMENT '出价金额USD',
  `amount_unit_usd` bigint(20) DEFAULT NULL COMMENT '出价金额8位',
  `bid_count` bigint(20) DEFAULT NULL COMMENT '出价总数',
  `min_bidder` varchar(50) DEFAULT NULL COMMENT '上一个最高出价值地址（当前出价起码要超过的最小金额地址）',
  `min_bid_unit_usd` bigint(20) DEFAULT NULL COMMENT '上一个最高出价值（当前出价起码要超过的最小金额数）',
  `payment_token` varchar(42) DEFAULT NULL COMMENT '支付代币地址',
  `transaction_hash` varchar(66) DEFAULT NULL COMMENT '交易哈希',
  `block_number` bigint(20) unsigned DEFAULT NULL COMMENT '区块号',
  `timestamp` bigint(20) unsigned DEFAULT NULL COMMENT '链上时间',
  `is_highest` tinyint(1) DEFAULT 0 COMMENT '是否为最高出价',
  `created_at` datetime NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_bids_user_id` (`user_id`),
  KEY `idx_bids_transaction_hash` (`transaction_hash`),
  KEY `idx_bids_is_highest` (`is_highest`),
  KEY `idx_bids_created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='出价表';

-- 数据导出被取消选择。

-- 导出  表 auction_market_db.nfts 结构
CREATE TABLE IF NOT EXISTS `nfts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `nft_id` varchar(64) NOT NULL COMMENT 'NFT唯一标识',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `contract_address` varchar(42) NOT NULL COMMENT 'NFT合约地址',
  `contract_name` varchar(255) DEFAULT NULL COMMENT '合约名称',
  `contract_symbol` varchar(64) DEFAULT NULL COMMENT '合约符号',
  `nft_owner_address` varchar(42) DEFAULT NULL COMMENT '当前拥有者地址(不是合约的要拥有者)',
  `token_id` bigint(20) unsigned DEFAULT NULL COMMENT 'Token ID',
  `token_uri` text DEFAULT NULL COMMENT 'Token URI',
  `nft_name` varchar(255) DEFAULT NULL COMMENT 'NFT名称',
  `image` text DEFAULT NULL COMMENT 'NFT图片URL',
  `description` text DEFAULT NULL COMMENT 'NFT描述',
  `metadata` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '完整元数据JSON' CHECK (json_valid(`metadata`)),
  `last_synced_at` datetime DEFAULT NULL COMMENT '上次同步时间',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_nft_id` (`nft_id`),
  KEY `idx_nfts_user_id` (`user_id`),
  KEY `idx_nfts_contract` (`contract_address`),
  KEY `idx_nfts_token_id` (`token_id`),
  KEY `idx_nfts_owner` (`nft_owner_address`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='NFT信息表';

-- 数据导出被取消选择。

-- 导出  表 auction_market_db.nft_ownerships 结构
CREATE TABLE IF NOT EXISTS `nft_ownerships` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nft_id` varchar(64) DEFAULT NULL COMMENT 'NFTID',
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户ID',
  `owner_address` varchar(50) DEFAULT NULL COMMENT 'NFT拥有者用户钱包地址（不一定是user_id对应的用户）',
  `status` varchar(50) DEFAULT NULL COMMENT '状态（holding,selling,sold,transfered）',
  `approved` tinyint(1) DEFAULT NULL COMMENT '是否授权给平台合约',
  `timestamp` bigint(20) DEFAULT NULL COMMENT '交易时间',
  `block_number` bigint(20) DEFAULT NULL COMMENT '区块数',
  `last_synced_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `nft_id_user_id` (`nft_id`,`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='nft用户关系表';

-- 数据导出被取消选择。

-- 导出  表 auction_market_db.users 结构
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `email` varchar(128) NOT NULL COMMENT '邮箱',
  `password` varchar(255) NOT NULL COMMENT '密码哈希',
  `wallet_address` varchar(42) DEFAULT NULL COMMENT '钱包地址',
  `nonce` varchar(64) DEFAULT NULL COMMENT '登录Nonce',
  `created_at` datetime NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_users_username` (`username`),
  UNIQUE KEY `idx_users_email` (`email`),
  KEY `idx_users_wallet_address` (`wallet_address`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 数据导出被取消选择。

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
