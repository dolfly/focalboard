// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

// Code generated by "make generate" from the Store interface
// DO NOT EDIT

// To add a public method, create an entry in the Store interface,
// prefix it with a @withTransaction comment if you need it to be
// transactional and then add a private method in the store itself
// with db sq.BaseRunner as the first parameter before running `make
// generate`

package sqlstore

import (
	"context"
	"time"

	"github.com/mattermost/focalboard/server/model"

	mmModel "github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

func (s *SQLStore) AddUpdateCategoryBoard(userID string, boardCategoryMapping map[string]string) error {
	if s.dbType == model.SqliteDBType {
		return s.addUpdateCategoryBoard(s.db, userID, boardCategoryMapping)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.addUpdateCategoryBoard(tx, userID, boardCategoryMapping)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "AddUpdateCategoryBoard"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) CanSeeUser(seerID string, seenID string) (bool, error) {
	return s.canSeeUser(s.db, seerID, seenID)

}

func (s *SQLStore) CleanUpSessions(expireTime int64) error {
	return s.cleanUpSessions(s.db, expireTime)

}

func (s *SQLStore) CreateBoardsAndBlocks(bab *model.BoardsAndBlocks, userID string) (*model.BoardsAndBlocks, error) {
	if s.dbType == model.SqliteDBType {
		return s.createBoardsAndBlocks(s.db, bab, userID)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return nil, txErr
	}
	result, err := s.createBoardsAndBlocks(tx, bab, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "CreateBoardsAndBlocks"))
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *SQLStore) CreateBoardsAndBlocksWithAdmin(bab *model.BoardsAndBlocks, userID string) (*model.BoardsAndBlocks, []*model.BoardMember, error) {
	if s.dbType == model.SqliteDBType {
		return s.createBoardsAndBlocksWithAdmin(s.db, bab, userID)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return nil, nil, txErr
	}
	result, resultVar1, err := s.createBoardsAndBlocksWithAdmin(tx, bab, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "CreateBoardsAndBlocksWithAdmin"))
		}
		return nil, nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, nil, err
	}

	return result, resultVar1, nil

}

func (s *SQLStore) CreateCategory(category model.Category) error {
	if s.dbType == model.SqliteDBType {
		return s.createCategory(s.db, category)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.createCategory(tx, category)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "CreateCategory"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) CreateSession(session *model.Session) error {
	return s.createSession(s.db, session)

}

func (s *SQLStore) CreateSubscription(sub *model.Subscription) (*model.Subscription, error) {
	return s.createSubscription(s.db, sub)

}

func (s *SQLStore) CreateUser(user *model.User) (*model.User, error) {
	return s.createUser(s.db, user)

}

func (s *SQLStore) DeleteBlock(blockID string, modifiedBy string) error {
	if s.dbType == model.SqliteDBType {
		return s.deleteBlock(s.db, blockID, modifiedBy)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.deleteBlock(tx, blockID, modifiedBy)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "DeleteBlock"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) DeleteBoard(boardID string, userID string) error {
	if s.dbType == model.SqliteDBType {
		return s.deleteBoard(s.db, boardID, userID)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.deleteBoard(tx, boardID, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "DeleteBoard"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) DeleteBoardsAndBlocks(dbab *model.DeleteBoardsAndBlocks, userID string) error {
	if s.dbType == model.SqliteDBType {
		return s.deleteBoardsAndBlocks(s.db, dbab, userID)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.deleteBoardsAndBlocks(tx, dbab, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "DeleteBoardsAndBlocks"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) DeleteCategory(categoryID string, userID string, teamID string) error {
	return s.deleteCategory(s.db, categoryID, userID, teamID)

}

func (s *SQLStore) DeleteMember(boardID string, userID string) error {
	return s.deleteMember(s.db, boardID, userID)

}

func (s *SQLStore) DeleteNotificationHint(blockID string) error {
	return s.deleteNotificationHint(s.db, blockID)

}

func (s *SQLStore) DeleteSession(sessionID string) error {
	return s.deleteSession(s.db, sessionID)

}

func (s *SQLStore) DeleteSubscription(blockID string, subscriberID string) error {
	return s.deleteSubscription(s.db, blockID, subscriberID)

}

func (s *SQLStore) DuplicateBlock(boardID string, blockID string, userID string, asTemplate bool) ([]*model.Block, error) {
	if s.dbType == model.SqliteDBType {
		return s.duplicateBlock(s.db, boardID, blockID, userID, asTemplate)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return nil, txErr
	}
	result, err := s.duplicateBlock(tx, boardID, blockID, userID, asTemplate)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "DuplicateBlock"))
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *SQLStore) DuplicateBoard(boardID string, userID string, toTeam string, asTemplate bool) (*model.BoardsAndBlocks, []*model.BoardMember, error) {
	if s.dbType == model.SqliteDBType {
		return s.duplicateBoard(s.db, boardID, userID, toTeam, asTemplate)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return nil, nil, txErr
	}
	result, resultVar1, err := s.duplicateBoard(tx, boardID, userID, toTeam, asTemplate)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "DuplicateBoard"))
		}
		return nil, nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, nil, err
	}

	return result, resultVar1, nil

}

func (s *SQLStore) GetActiveUserCount(updatedSecondsAgo int64) (int, error) {
	return s.getActiveUserCount(s.db, updatedSecondsAgo)

}

func (s *SQLStore) GetAllTeams() ([]*model.Team, error) {
	return s.getAllTeams(s.db)

}

func (s *SQLStore) GetBlock(blockID string) (*model.Block, error) {
	return s.getBlock(s.db, blockID)

}

func (s *SQLStore) GetBlockCountsByType() (map[string]int64, error) {
	return s.getBlockCountsByType(s.db)

}

func (s *SQLStore) GetBlockHistory(blockID string, opts model.QueryBlockHistoryOptions) ([]*model.Block, error) {
	return s.getBlockHistory(s.db, blockID, opts)

}

func (s *SQLStore) GetBlockHistoryDescendants(boardID string, opts model.QueryBlockHistoryOptions) ([]*model.Block, error) {
	return s.getBlockHistoryDescendants(s.db, boardID, opts)

}

func (s *SQLStore) GetBlocks(opts model.QueryBlocksOptions) ([]*model.Block, error) {
	return s.getBlocks(s.db, opts)

}

func (s *SQLStore) GetBlocksByIDs(ids []string) ([]*model.Block, error) {
	return s.getBlocksByIDs(s.db, ids)

}

func (s *SQLStore) GetBlocksForBoard(boardID string) ([]*model.Block, error) {
	return s.getBlocksForBoard(s.db, boardID)

}

func (s *SQLStore) GetBlocksWithParent(boardID string, parentID string) ([]*model.Block, error) {
	return s.getBlocksWithParent(s.db, boardID, parentID)

}

func (s *SQLStore) GetBlocksWithParentAndType(boardID string, parentID string, blockType string) ([]*model.Block, error) {
	return s.getBlocksWithParentAndType(s.db, boardID, parentID, blockType)

}

func (s *SQLStore) GetBlocksWithType(boardID string, blockType string) ([]*model.Block, error) {
	return s.getBlocksWithType(s.db, boardID, blockType)

}

func (s *SQLStore) GetBoard(id string) (*model.Board, error) {
	return s.getBoard(s.db, id)

}

func (s *SQLStore) GetBoardAndCard(block *model.Block) (*model.Board, *model.Block, error) {
	return s.getBoardAndCard(s.db, block)

}

func (s *SQLStore) GetBoardAndCardByID(blockID string) (*model.Board, *model.Block, error) {
	return s.getBoardAndCardByID(s.db, blockID)

}

func (s *SQLStore) GetBoardCount() (int64, error) {
	return s.getBoardCount(s.db)

}

func (s *SQLStore) GetBoardHistory(boardID string, opts model.QueryBoardHistoryOptions) ([]*model.Board, error) {
	return s.getBoardHistory(s.db, boardID, opts)

}

func (s *SQLStore) GetBoardMemberHistory(boardID string, userID string, limit uint64) ([]*model.BoardMemberHistoryEntry, error) {
	return s.getBoardMemberHistory(s.db, boardID, userID, limit)

}

func (s *SQLStore) GetBoardsForUserAndTeam(userID string, teamID string, includePublicBoards bool) ([]*model.Board, error) {
	return s.getBoardsForUserAndTeam(s.db, userID, teamID, includePublicBoards)

}

func (s *SQLStore) GetBoardsInTeamByIds(boardIDs []string, teamID string) ([]*model.Board, error) {
	return s.getBoardsInTeamByIds(s.db, boardIDs, teamID)

}

func (s *SQLStore) GetCardLimitTimestamp() (int64, error) {
	return s.getCardLimitTimestamp(s.db)

}

func (s *SQLStore) GetCategory(id string) (*model.Category, error) {
	return s.getCategory(s.db, id)

}

func (s *SQLStore) GetChannel(teamID string, channelID string) (*mmModel.Channel, error) {
	return s.getChannel(s.db, teamID, channelID)

}

func (s *SQLStore) GetCloudLimits() (*mmModel.ProductLimits, error) {
	return s.getCloudLimits(s.db)

}

func (s *SQLStore) GetFileInfo(id string) (*mmModel.FileInfo, error) {
	return s.getFileInfo(s.db, id)

}

func (s *SQLStore) GetLicense() *mmModel.License {
	return s.getLicense(s.db)

}

func (s *SQLStore) GetMemberForBoard(boardID string, userID string) (*model.BoardMember, error) {
	return s.getMemberForBoard(s.db, boardID, userID)

}

func (s *SQLStore) GetMembersForBoard(boardID string) ([]*model.BoardMember, error) {
	return s.getMembersForBoard(s.db, boardID)

}

func (s *SQLStore) GetMembersForUser(userID string) ([]*model.BoardMember, error) {
	return s.getMembersForUser(s.db, userID)

}

func (s *SQLStore) GetNextNotificationHint(remove bool) (*model.NotificationHint, error) {
	return s.getNextNotificationHint(s.db, remove)

}

func (s *SQLStore) GetNotificationHint(blockID string) (*model.NotificationHint, error) {
	return s.getNotificationHint(s.db, blockID)

}

func (s *SQLStore) GetRegisteredUserCount() (int, error) {
	return s.getRegisteredUserCount(s.db)

}

func (s *SQLStore) GetSession(token string, expireTime int64) (*model.Session, error) {
	return s.getSession(s.db, token, expireTime)

}

func (s *SQLStore) GetSharing(rootID string) (*model.Sharing, error) {
	return s.getSharing(s.db, rootID)

}

func (s *SQLStore) GetSubTree2(boardID string, blockID string, opts model.QuerySubtreeOptions) ([]*model.Block, error) {
	return s.getSubTree2(s.db, boardID, blockID, opts)

}

func (s *SQLStore) GetSubscribersCountForBlock(blockID string) (int, error) {
	return s.getSubscribersCountForBlock(s.db, blockID)

}

func (s *SQLStore) GetSubscribersForBlock(blockID string) ([]*model.Subscriber, error) {
	return s.getSubscribersForBlock(s.db, blockID)

}

func (s *SQLStore) GetSubscription(blockID string, subscriberID string) (*model.Subscription, error) {
	return s.getSubscription(s.db, blockID, subscriberID)

}

func (s *SQLStore) GetSubscriptions(subscriberID string) ([]*model.Subscription, error) {
	return s.getSubscriptions(s.db, subscriberID)

}

func (s *SQLStore) GetSystemSetting(key string) (string, error) {
	return s.getSystemSetting(s.db, key)

}

func (s *SQLStore) GetSystemSettings() (map[string]string, error) {
	return s.getSystemSettings(s.db)

}

func (s *SQLStore) GetTeam(ID string) (*model.Team, error) {
	return s.getTeam(s.db, ID)

}

func (s *SQLStore) GetTeamBoardsInsights(teamID string, userID string, since int64, offset int, limit int, boardIDs []string) (*model.BoardInsightsList, error) {
	return s.getTeamBoardsInsights(s.db, teamID, userID, since, offset, limit, boardIDs)

}

func (s *SQLStore) GetTeamCount() (int64, error) {
	return s.getTeamCount(s.db)

}

func (s *SQLStore) GetTeamsForUser(userID string) ([]*model.Team, error) {
	return s.getTeamsForUser(s.db, userID)

}

func (s *SQLStore) GetTemplateBoards(teamID string, userID string) ([]*model.Board, error) {
	return s.getTemplateBoards(s.db, teamID, userID)

}

func (s *SQLStore) GetUsedCardsCount() (int, error) {
	return s.getUsedCardsCount(s.db)

}

func (s *SQLStore) GetUserBoardsInsights(teamID string, userID string, since int64, offset int, limit int, boardIDs []string) (*model.BoardInsightsList, error) {
	return s.getUserBoardsInsights(s.db, teamID, userID, since, offset, limit, boardIDs)

}

func (s *SQLStore) GetUserByEmail(email string) (*model.User, error) {
	return s.getUserByEmail(s.db, email)

}

func (s *SQLStore) GetUserByID(userID string) (*model.User, error) {
	return s.getUserByID(s.db, userID)

}

func (s *SQLStore) GetUserByUsername(username string) (*model.User, error) {
	return s.getUserByUsername(s.db, username)

}

func (s *SQLStore) GetUserCategories(userID string, teamID string) ([]model.Category, error) {
	return s.getUserCategories(s.db, userID, teamID)

}

func (s *SQLStore) GetUserCategoryBoards(userID string, teamID string) ([]model.CategoryBoards, error) {
	return s.getUserCategoryBoards(s.db, userID, teamID)

}

func (s *SQLStore) GetUserPreferences(userID string) (mmModel.Preferences, error) {
	return s.getUserPreferences(s.db, userID)

}

func (s *SQLStore) GetUserTimezone(userID string) (string, error) {
	return s.getUserTimezone(s.db, userID)

}

func (s *SQLStore) GetUsersByTeam(teamID string, asGuestID string, showEmail bool, showName bool) ([]*model.User, error) {
	return s.getUsersByTeam(s.db, teamID, asGuestID, showEmail, showName)

}

func (s *SQLStore) GetUsersList(userIDs []string, showEmail bool, showName bool) ([]*model.User, error) {
	return s.getUsersList(s.db, userIDs, showEmail, showName)

}

func (s *SQLStore) InsertBlock(block *model.Block, userID string) error {
	if s.dbType == model.SqliteDBType {
		return s.insertBlock(s.db, block, userID)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.insertBlock(tx, block, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "InsertBlock"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) InsertBlocks(blocks []*model.Block, userID string) error {
	if s.dbType == model.SqliteDBType {
		return s.insertBlocks(s.db, blocks, userID)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.insertBlocks(tx, blocks, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "InsertBlocks"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) InsertBoard(board *model.Board, userID string) (*model.Board, error) {
	return s.insertBoard(s.db, board, userID)

}

func (s *SQLStore) InsertBoardWithAdmin(board *model.Board, userID string) (*model.Board, *model.BoardMember, error) {
	if s.dbType == model.SqliteDBType {
		return s.insertBoardWithAdmin(s.db, board, userID)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return nil, nil, txErr
	}
	result, resultVar1, err := s.insertBoardWithAdmin(tx, board, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "InsertBoardWithAdmin"))
		}
		return nil, nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, nil, err
	}

	return result, resultVar1, nil

}

func (s *SQLStore) PatchBlock(blockID string, blockPatch *model.BlockPatch, userID string) error {
	if s.dbType == model.SqliteDBType {
		return s.patchBlock(s.db, blockID, blockPatch, userID)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.patchBlock(tx, blockID, blockPatch, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "PatchBlock"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) PatchBlocks(blockPatches *model.BlockPatchBatch, userID string) error {
	if s.dbType == model.SqliteDBType {
		return s.patchBlocks(s.db, blockPatches, userID)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.patchBlocks(tx, blockPatches, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "PatchBlocks"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) PatchBoard(boardID string, boardPatch *model.BoardPatch, userID string) (*model.Board, error) {
	if s.dbType == model.SqliteDBType {
		return s.patchBoard(s.db, boardID, boardPatch, userID)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return nil, txErr
	}
	result, err := s.patchBoard(tx, boardID, boardPatch, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "PatchBoard"))
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *SQLStore) PatchBoardsAndBlocks(pbab *model.PatchBoardsAndBlocks, userID string) (*model.BoardsAndBlocks, error) {
	if s.dbType == model.SqliteDBType {
		return s.patchBoardsAndBlocks(s.db, pbab, userID)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return nil, txErr
	}
	result, err := s.patchBoardsAndBlocks(tx, pbab, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "PatchBoardsAndBlocks"))
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *SQLStore) PatchUserPreferences(userID string, patch model.UserPreferencesPatch) (mmModel.Preferences, error) {
	return s.patchUserPreferences(s.db, userID, patch)

}

func (s *SQLStore) PostMessage(message string, postType string, channelID string) error {
	return s.postMessage(s.db, message, postType, channelID)

}

func (s *SQLStore) RefreshSession(session *model.Session) error {
	return s.refreshSession(s.db, session)

}

func (s *SQLStore) RemoveDefaultTemplates(boards []*model.Board) error {
	return s.removeDefaultTemplates(s.db, boards)

}

func (s *SQLStore) ReorderCategories(userID string, teamID string, newCategoryOrder []string) ([]string, error) {
	return s.reorderCategories(s.db, userID, teamID, newCategoryOrder)

}

func (s *SQLStore) ReorderCategoryBoards(categoryID string, newBoardsOrder []string) ([]string, error) {
	return s.reorderCategoryBoards(s.db, categoryID, newBoardsOrder)

}

func (s *SQLStore) RunDataRetention(globalRetentionDate int64, batchSize int64) (int64, error) {
	if s.dbType == model.SqliteDBType {
		return s.runDataRetention(s.db, globalRetentionDate, batchSize)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return 0, txErr
	}
	result, err := s.runDataRetention(tx, globalRetentionDate, batchSize)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "RunDataRetention"))
		}
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return result, nil

}

func (s *SQLStore) SaveFileInfo(fileInfo *mmModel.FileInfo) error {
	return s.saveFileInfo(s.db, fileInfo)

}

func (s *SQLStore) SaveMember(bm *model.BoardMember) (*model.BoardMember, error) {
	return s.saveMember(s.db, bm)

}

func (s *SQLStore) SearchBoardsForUser(term string, userID string, includePublicBoards bool) ([]*model.Board, error) {
	return s.searchBoardsForUser(s.db, term, userID, includePublicBoards)

}

func (s *SQLStore) SearchBoardsForUserInTeam(teamID string, term string, userID string) ([]*model.Board, error) {
	return s.searchBoardsForUserInTeam(s.db, teamID, term, userID)

}

func (s *SQLStore) SearchUserChannels(teamID string, userID string, query string) ([]*mmModel.Channel, error) {
	return s.searchUserChannels(s.db, teamID, userID, query)

}

func (s *SQLStore) SearchUsersByTeam(teamID string, searchQuery string, asGuestID string, excludeBots bool, showEmail bool, showName bool) ([]*model.User, error) {
	return s.searchUsersByTeam(s.db, teamID, searchQuery, asGuestID, excludeBots, showEmail, showName)

}

func (s *SQLStore) SendMessage(message string, postType string, receipts []string) error {
	return s.sendMessage(s.db, message, postType, receipts)

}

func (s *SQLStore) SetSystemSetting(key string, value string) error {
	return s.setSystemSetting(s.db, key, value)

}

func (s *SQLStore) UndeleteBlock(blockID string, modifiedBy string) error {
	if s.dbType == model.SqliteDBType {
		return s.undeleteBlock(s.db, blockID, modifiedBy)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.undeleteBlock(tx, blockID, modifiedBy)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "UndeleteBlock"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) UndeleteBoard(boardID string, modifiedBy string) error {
	if s.dbType == model.SqliteDBType {
		return s.undeleteBoard(s.db, boardID, modifiedBy)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.undeleteBoard(tx, boardID, modifiedBy)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "UndeleteBoard"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) UpdateCardLimitTimestamp(cardLimit int) (int64, error) {
	return s.updateCardLimitTimestamp(s.db, cardLimit)

}

func (s *SQLStore) UpdateCategory(category model.Category) error {
	return s.updateCategory(s.db, category)

}

func (s *SQLStore) UpdateSession(session *model.Session) error {
	return s.updateSession(s.db, session)

}

func (s *SQLStore) UpdateSubscribersNotifiedAt(blockID string, notifiedAt int64) error {
	return s.updateSubscribersNotifiedAt(s.db, blockID, notifiedAt)

}

func (s *SQLStore) UpdateUser(user *model.User) (*model.User, error) {
	return s.updateUser(s.db, user)

}

func (s *SQLStore) UpdateUserPassword(username string, password string) error {
	return s.updateUserPassword(s.db, username, password)

}

func (s *SQLStore) UpdateUserPasswordByID(userID string, password string) error {
	return s.updateUserPasswordByID(s.db, userID, password)

}

func (s *SQLStore) UpsertNotificationHint(hint *model.NotificationHint, notificationFreq time.Duration) (*model.NotificationHint, error) {
	return s.upsertNotificationHint(s.db, hint, notificationFreq)

}

func (s *SQLStore) UpsertSharing(sharing model.Sharing) error {
	return s.upsertSharing(s.db, sharing)

}

func (s *SQLStore) UpsertTeamSettings(team model.Team) error {
	return s.upsertTeamSettings(s.db, team)

}

func (s *SQLStore) UpsertTeamSignupToken(team model.Team) error {
	return s.upsertTeamSignupToken(s.db, team)

}
