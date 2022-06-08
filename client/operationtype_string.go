// Code generated by "stringer -type=OperationType"; DO NOT EDIT.

package client

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[opUnused-0]
	_ = x[opPing-1]
	_ = x[opJoin-2]
	_ = x[opCreateAccount-3]
	_ = x[opLogin-4]
	_ = x[opSendCrashLog-5]
	_ = x[opSendTraceRoute-6]
	_ = x[opSendVfxStats-7]
	_ = x[opSendGamePingInfo-8]
	_ = x[opCreateCharacter-9]
	_ = x[opDeleteCharacter-10]
	_ = x[opSelectCharacter-11]
	_ = x[opRedeemKeycode-12]
	_ = x[opGetGameServerByCluster-13]
	_ = x[opGetActiveSubscription-14]
	_ = x[opGetShopPurchaseUrl-15]
	_ = x[opGetBuyTrialDetails-16]
	_ = x[opGetReferralSeasonDetails-17]
	_ = x[opGetReferralLink-18]
	_ = x[opGetAvailableTrialKeys-19]
	_ = x[opGetShopTilesForCategory-20]
	_ = x[opMove-21]
	_ = x[opCastStart-22]
	_ = x[opCastCancel-23]
	_ = x[opTerminateToggleSpell-24]
	_ = x[opChannelingCancel-25]
	_ = x[opAttackBuildingStart-26]
	_ = x[opInventoryDestroyItem-27]
	_ = x[opInventoryMoveItem-28]
	_ = x[opInventoryRecoverItem-29]
	_ = x[opInventoryRecoverAllItems-30]
	_ = x[opInventorySplitStack-31]
	_ = x[opInventorySplitStackInto-32]
	_ = x[opGetClusterData-33]
	_ = x[opChangeCluster-34]
	_ = x[opConsoleCommand-35]
	_ = x[opChatMessage-36]
	_ = x[opReportClientError-37]
	_ = x[opRegisterToObject-38]
	_ = x[opUnRegisterFromObject-39]
	_ = x[opCraftBuildingChangeSettings-40]
	_ = x[opCraftBuildingTakeMoney-41]
	_ = x[opRepairBuildingChangeSettings-42]
	_ = x[opRepairBuildingTakeMoney-43]
	_ = x[opActionBuildingChangeSettings-44]
	_ = x[opHarvestStart-45]
	_ = x[opHarvestCancel-46]
	_ = x[opTakeSilver-47]
	_ = x[opActionOnBuildingStart-48]
	_ = x[opActionOnBuildingCancel-49]
	_ = x[opItemRerollQualityStart-50]
	_ = x[opItemRerollQualityCancel-51]
	_ = x[opInstallResourceStart-52]
	_ = x[opInstallResourceCancel-53]
	_ = x[opInstallSilver-54]
	_ = x[opBuildingFillNutrition-55]
	_ = x[opBuildingChangeRenovationState-56]
	_ = x[opBuildingBuySkin-57]
	_ = x[opBuildingClaim-58]
	_ = x[opBuildingGiveup-59]
	_ = x[opBuildingNutritionSilverStorageDeposit-60]
	_ = x[opBuildingNutritionSilverStorageWithdraw-61]
	_ = x[opBuildingNutritionSilverRewardSet-62]
	_ = x[opConstructionSiteCreate-63]
	_ = x[opPlaceableObjectPlace-64]
	_ = x[opPlaceableObjectPlaceCancel-65]
	_ = x[opPlaceableObjectPickup-66]
	_ = x[opFurnitureObjectUse-67]
	_ = x[opFarmableHarvest-68]
	_ = x[opFarmableFinishGrownItem-69]
	_ = x[opFarmableDestroy-70]
	_ = x[opFarmableGetProduct-71]
	_ = x[opTearDownConstructionSite-72]
	_ = x[opCastleGateUse-73]
	_ = x[opAuctionCreateRequest-74]
	_ = x[opAuctionCreateOffer-75]
	_ = x[opAuctionGetOffers-76]
	_ = x[opAuctionGetRequests-77]
	_ = x[opAuctionBuyOffer-78]
	_ = x[opAuctionAbortAuction-79]
	_ = x[opAuctionModifyAuction-80]
	_ = x[opAuctionAbortOffer-81]
	_ = x[opAuctionAbortRequest-82]
	_ = x[opAuctionSellRequest-83]
	_ = x[opAuctionGetFinishedAuctions-84]
	_ = x[opAuctionGetFinishedAuctionsCount-85]
	_ = x[opAuctionFetchAuction-86]
	_ = x[opAuctionGetMyOpenOffers-87]
	_ = x[opAuctionGetMyOpenRequests-88]
	_ = x[opAuctionGetMyOpenAuctions-89]
	_ = x[opAuctionGetItemAverageStats-90]
	_ = x[opAuctionGetItemAverageValue-91]
	_ = x[opContainerOpen-92]
	_ = x[opContainerClose-93]
	_ = x[opContainerManageSubContainer-94]
	_ = x[opRespawn-95]
	_ = x[opSuicide-96]
	_ = x[opJoinGuild-97]
	_ = x[opLeaveGuild-98]
	_ = x[opCreateGuild-99]
	_ = x[opInviteToGuild-100]
	_ = x[opDeclineGuildInvitation-101]
	_ = x[opKickFromGuild-102]
	_ = x[opDuellingChallengePlayer-103]
	_ = x[opDuellingAcceptChallenge-104]
	_ = x[opDuellingDenyChallenge-105]
	_ = x[opChangeClusterTax-106]
	_ = x[opClaimTerritory-107]
	_ = x[opGiveUpTerritory-108]
	_ = x[opChangeTerritoryAccessRights-109]
	_ = x[opGetMonolithInfo-110]
	_ = x[opGetClaimInfo-111]
	_ = x[opGetAttackInfo-112]
	_ = x[opGetTerritorySeasonPoints-113]
	_ = x[opGetAttackSchedule-114]
	_ = x[opScheduleAttack-115]
	_ = x[opGetMatches-116]
	_ = x[opGetMatchDetails-117]
	_ = x[opJoinMatch-118]
	_ = x[opLeaveMatch-119]
	_ = x[opChangeChatSettings-120]
	_ = x[opLogoutStart-121]
	_ = x[opLogoutCancel-122]
	_ = x[opClaimOrbStart-123]
	_ = x[opClaimOrbCancel-124]
	_ = x[opMatchLootChestOpeningStart-125]
	_ = x[opMatchLootChestOpeningCancel-126]
	_ = x[opDepositToGuildAccount-127]
	_ = x[opWithdrawalFromAccount-128]
	_ = x[opChangeGuildPayUpkeepFlag-129]
	_ = x[opChangeGuildTax-130]
	_ = x[opGetMyTerritories-131]
	_ = x[opMorganaCommand-132]
	_ = x[opGetServerInfo-133]
	_ = x[opInviteMercenaryToMatch-134]
	_ = x[opSubscribeToCluster-135]
	_ = x[opAnswerMercenaryInvitation-136]
	_ = x[opGetCharacterEquipment-137]
	_ = x[opGetCharacterSteamAchievements-138]
	_ = x[opGetCharacterStats-139]
	_ = x[opGetKillHistoryDetails-140]
	_ = x[opLearnMasteryLevel-141]
	_ = x[opReSpecAchievement-142]
	_ = x[opChangeAvatar-143]
	_ = x[opGetRankings-144]
	_ = x[opGetRank-145]
	_ = x[opGetGvgSeasonRankings-146]
	_ = x[opGetGvgSeasonRank-147]
	_ = x[opGetGvgSeasonHistoryRankings-148]
	_ = x[opGetGvgSeasonGuildMemberHistory-149]
	_ = x[opKickFromGvGMatch-150]
	_ = x[opGetChestLogs-151]
	_ = x[opGetAccessRightLogs-152]
	_ = x[opGetGuildAccountLogs-153]
	_ = x[opGetGuildAccountLogsLargeAmount-154]
	_ = x[opInviteToPlayerTrade-155]
	_ = x[opPlayerTradeCancel-156]
	_ = x[opPlayerTradeInvitationAccept-157]
	_ = x[opPlayerTradeAddItem-158]
	_ = x[opPlayerTradeRemoveItem-159]
	_ = x[opPlayerTradeAcceptTrade-160]
	_ = x[opPlayerTradeSetSilverOrGold-161]
	_ = x[opSendMiniMapPing-162]
	_ = x[opStuck-163]
	_ = x[opBuyRealEstate-164]
	_ = x[opClaimRealEstate-165]
	_ = x[opGiveUpRealEstate-166]
	_ = x[opChangeRealEstateOutline-167]
	_ = x[opGetMailInfos-168]
	_ = x[opGetMailCount-169]
	_ = x[opReadMail-170]
	_ = x[opSendNewMail-171]
	_ = x[opDeleteMail-172]
	_ = x[opMarkMailUnread-173]
	_ = x[opClaimAttachmentFromMail-174]
	_ = x[opUpdateLfgInfo-175]
	_ = x[opGetLfgInfos-176]
	_ = x[opGetMyGuildLfgInfo-177]
	_ = x[opGetLfgDescriptionText-178]
	_ = x[opLfgApplyToGuild-179]
	_ = x[opAnswerLfgGuildApplication-180]
	_ = x[opRegisterChatPeer-181]
	_ = x[opSendChatMessage-182]
	_ = x[opJoinChatChannel-183]
	_ = x[opLeaveChatChannel-184]
	_ = x[opSendWhisperMessage-185]
	_ = x[opSay-186]
	_ = x[opPlayEmote-187]
	_ = x[opStopEmote-188]
	_ = x[opGetClusterMapInfo-189]
	_ = x[opAccessRightsChangeSettings-190]
	_ = x[opMount-191]
	_ = x[opMountCancel-192]
	_ = x[opBuyJourney-193]
	_ = x[opSetSaleStatusForEstate-194]
	_ = x[opResolveGuildOrPlayerName-195]
	_ = x[opGetRespawnInfos-196]
	_ = x[opMakeHome-197]
	_ = x[opLeaveHome-198]
	_ = x[opResurrectionReply-199]
	_ = x[opAllianceCreate-200]
	_ = x[opAllianceDisband-201]
	_ = x[opAllianceGetMemberInfos-202]
	_ = x[opAllianceInvite-203]
	_ = x[opAllianceAnswerInvitation-204]
	_ = x[opAllianceCancelInvitation-205]
	_ = x[opAllianceKickGuild-206]
	_ = x[opAllianceLeave-207]
	_ = x[opAllianceChangeGoldPaymentFlag-208]
	_ = x[opAllianceGetDetailInfo-209]
	_ = x[opGetIslandInfos-210]
	_ = x[opAbandonMyIsland-211]
	_ = x[opBuyMyIsland-212]
	_ = x[opBuyGuildIsland-213]
	_ = x[opAbandonGuildIsland-214]
	_ = x[opUpgradeMyIsland-215]
	_ = x[opUpgradeGuildIsland-216]
	_ = x[opMoveMyIsland-217]
	_ = x[opMoveGuildIsland-218]
	_ = x[opTerritoryFillNutrition-219]
	_ = x[opTeleportBack-220]
	_ = x[opPartyInvitePlayer-221]
	_ = x[opPartyAnswerInvitation-222]
	_ = x[opPartyLeave-223]
	_ = x[opPartyKickPlayer-224]
	_ = x[opPartyMakeLeader-225]
	_ = x[opPartyChangeLootSetting-226]
	_ = x[opPartyMarkObject-227]
	_ = x[opPartySetRole-228]
	_ = x[opGetGuildMOTD-229]
	_ = x[opSetGuildMOTD-230]
	_ = x[opExitEnterStart-231]
	_ = x[opExitEnterCancel-232]
	_ = x[opQuestGiverRequest-233]
	_ = x[opGoldMarketGetBuyOffer-234]
	_ = x[opGoldMarketGetBuyOfferFromSilver-235]
	_ = x[opGoldMarketGetSellOffer-236]
	_ = x[opGoldMarketGetSellOfferFromSilver-237]
	_ = x[opGoldMarketBuyGold-238]
	_ = x[opGoldMarketSellGold-239]
	_ = x[opGoldMarketCreateSellOrder-240]
	_ = x[opGoldMarketCreateBuyOrder-241]
	_ = x[opGoldMarketGetInfos-242]
	_ = x[opGoldMarketCancelOrder-243]
	_ = x[opUnknown244-244]
	_ = x[opUnknown245-245]
	_ = x[opUnknown246-246]
	_ = x[opGoldMarketGetAverageInfo-247]
	_ = x[opSiegeCampClaimStart-248]
	_ = x[opSiegeCampClaimCancel-249]
	_ = x[opTreasureChestUsingStart-250]
	_ = x[opTreasureChestUsingCancel-251]
	_ = x[opUseLootChest-252]
	_ = x[opUseShrine-253]
	_ = x[opLaborerStartJob-254]
	_ = x[opLaborerTakeJobLoot-255]
	_ = x[opLaborerDismiss-256]
	_ = x[opLaborerMove-257]
	_ = x[opLaborerBuyItem-258]
	_ = x[opLaborerUpgrade-259]
	_ = x[opBuyPremium-260]
	_ = x[opBuyTrial-261]
	_ = x[opRealEstateGetAuctionData-262]
	_ = x[opRealEstateBidOnAuction-263]
	_ = x[opGetSiegeCampCooldown-264]
	_ = x[opFriendInvite-265]
	_ = x[opFriendAnswerInvitation-266]
	_ = x[opFriendCancelnvitation-267]
	_ = x[opFriendRemove-268]
	_ = x[opInventoryStack-269]
	_ = x[opInventorySort-270]
	_ = x[opEquipmentItemChangeSpell-271]
	_ = x[opExpeditionRegister-272]
	_ = x[opExpeditionRegisterCancel-273]
	_ = x[opJoinExpedition-274]
	_ = x[opDeclineExpeditionInvitation-275]
	_ = x[opVoteStart-276]
	_ = x[opVoteDoVote-277]
	_ = x[opRatingDoRate-278]
	_ = x[opEnteringExpeditionStart-279]
	_ = x[opEnteringExpeditionCancel-280]
	_ = x[opActivateExpeditionCheckPoint-281]
	_ = x[opArenaRegister-282]
	_ = x[opArenaRegisterCancel-283]
	_ = x[opArenaLeave-284]
	_ = x[opJoinArenaMatch-285]
	_ = x[opDeclineArenaInvitation-286]
	_ = x[opEnteringArenaStart-287]
	_ = x[opEnteringArenaCancel-288]
	_ = x[opArenaCustomMatch-289]
	_ = x[opArenaCustomMatchCreate-290]
	_ = x[opUpdateCharacterStatement-291]
	_ = x[opBoostFarmable-292]
	_ = x[opGetStrikeHistory-293]
	_ = x[opUseFunction-294]
	_ = x[opUsePortalEntrance-295]
	_ = x[opResetPortalBinding-296]
	_ = x[opQueryPortalBinding-297]
	_ = x[opClaimPaymentTransaction-298]
	_ = x[opChangeUseFlag-299]
	_ = x[opClientPerformanceStats-300]
	_ = x[opExtendedHardwareStats-301]
	_ = x[opClientLowMemoryWarning-302]
	_ = x[opTerritoryClaimStart-303]
	_ = x[opTerritoryClaimCancel-304]
	_ = x[opRequestAppStoreProducts-305]
	_ = x[opVerifyProductPurchase-306]
	_ = x[opQueryGuildPlayerStats-307]
	_ = x[opQueryAllianceGuildStats-308]
	_ = x[opTrackAchievements-309]
	_ = x[opSetAchievementsAutoLearn-310]
	_ = x[opDepositItemToGuildCurrency-311]
	_ = x[opWithdrawalItemFromGuildCurrency-312]
	_ = x[opAuctionSellSpecificItemRequest-313]
	_ = x[opFishingStart-314]
	_ = x[opFishingCasting-315]
	_ = x[opFishingCast-316]
	_ = x[opFishingCatch-317]
	_ = x[opFishingPull-318]
	_ = x[opFishingGiveLine-319]
	_ = x[opFishingFinish-320]
	_ = x[opFishingCancel-321]
	_ = x[opCreateGuildAccessTag-322]
	_ = x[opDeleteGuildAccessTag-323]
	_ = x[opRenameGuildAccessTag-324]
	_ = x[opFlagGuildAccessTagGuildPermission-325]
	_ = x[opAssignGuildAccessTag-326]
	_ = x[opRemoveGuildAccessTagFromPlayer-327]
	_ = x[opModifyGuildAccessTagEditors-328]
	_ = x[opRequestPublicAccessTags-329]
	_ = x[opChangeAccessTagPublicFlag-330]
	_ = x[opUpdateGuildAccessTag-331]
	_ = x[opSteamStartMicrotransaction-332]
	_ = x[opSteamFinishMicrotransaction-333]
	_ = x[opSteamIdHasActiveAccount-334]
	_ = x[opCheckEmailAccountState-335]
	_ = x[opLinkAccountToSteamId-336]
	_ = x[opBuyGvgSeasonBooster-337]
	_ = x[opChangeFlaggingPrepare-338]
	_ = x[opOverCharge-339]
	_ = x[opOverChargeEnd-340]
	_ = x[opRequestTrusted-341]
	_ = x[opChangeGuildLogo-342]
	_ = x[opPartyFinderRegisterForUpdates-343]
	_ = x[opPartyFinderUnregisterForUpdates-344]
	_ = x[opPartyFinderEnlistNewPartySearch-345]
	_ = x[opPartyFinderDeletePartySearch-346]
	_ = x[opPartyFinderChangePartySearch-347]
	_ = x[opPartyFinderChangeRole-348]
	_ = x[opPartyFinderApplyForGroup-349]
	_ = x[opPartyFinderAcceptOrDeclineApplyForGroup-350]
	_ = x[opPartyFinderGetEquipmentSnapshot-351]
	_ = x[opPartyFinderRegisterApplicants-352]
	_ = x[opPartyFinderUnregisterApplicants-353]
	_ = x[opPartyFinderFulltextSearch-354]
	_ = x[opPartyFinderRequestEquipmentSnapshot-355]
	_ = x[opGetPersonalSeasonTrackerData-356]
	_ = x[opUseConsumableFromInventory-357]
	_ = x[opClaimPersonalSeasonReward-358]
	_ = x[opEasyAntiCheatMessageToServer-359]
	_ = x[opSetNextTutorialState-360]
	_ = x[opAddPlayerToMuteList-361]
	_ = x[opRemovePlayerFromMuteList-362]
	_ = x[opProductShopUserEvent-363]
	_ = x[opGetVanityUnlocks-364]
	_ = x[opBuyVanityUnlocks-365]
	_ = x[opGetMountSkins-366]
	_ = x[opSetMountSkin-367]
	_ = x[opSetWardrobe-368]
	_ = x[opChangeCustomization-369]
	_ = x[opSetFavoriteIsland-370]
	_ = x[opGetGuildChallengePoints-371]
	_ = x[opTravelToHideout-372]
	_ = x[opSmartQueueJoin-373]
	_ = x[opSmartQueueLeave-374]
	_ = x[opSmartQueueSelectSpawnCluster-375]
	_ = x[opUpgradeHideout-376]
	_ = x[opInitHideoutAttackStart-377]
	_ = x[opInitHideoutAttackCancel-378]
	_ = x[opHideoutFillNutrition-379]
	_ = x[opHideoutGetInfo-380]
	_ = x[opHideoutGetOwnerInfo-381]
	_ = x[opHideoutSetTribute-382]
	_ = x[opOpenWorldAttackScheduleStart-383]
	_ = x[opOpenWorldAttackScheduleCancel-384]
	_ = x[opOpenWorldAttackConquerStart-385]
	_ = x[opOpenWorldAttackConquerCancel-386]
	_ = x[opGetOpenWorldAttackDetails-387]
	_ = x[opGetNextOpenWorldAttackScheduleTime-388]
	_ = x[opRecoverVaultFromHideout-389]
	_ = x[opGetGuildEnergyDrainInfo-390]
	_ = x[opChannelingUpdate-391]
}

const _OperationType_name = "opUnusedopPingopJoinopCreateAccountopLoginopSendCrashLogopSendTraceRouteopSendVfxStatsopSendGamePingInfoopCreateCharacteropDeleteCharacteropSelectCharacteropRedeemKeycodeopGetGameServerByClusteropGetActiveSubscriptionopGetShopPurchaseUrlopGetBuyTrialDetailsopGetReferralSeasonDetailsopGetReferralLinkopGetAvailableTrialKeysopGetShopTilesForCategoryopMoveopCastStartopCastCancelopTerminateToggleSpellopChannelingCancelopAttackBuildingStartopInventoryDestroyItemopInventoryMoveItemopInventoryRecoverItemopInventoryRecoverAllItemsopInventorySplitStackopInventorySplitStackIntoopGetClusterDataopChangeClusteropConsoleCommandopChatMessageopReportClientErroropRegisterToObjectopUnRegisterFromObjectopCraftBuildingChangeSettingsopCraftBuildingTakeMoneyopRepairBuildingChangeSettingsopRepairBuildingTakeMoneyopActionBuildingChangeSettingsopHarvestStartopHarvestCancelopTakeSilveropActionOnBuildingStartopActionOnBuildingCancelopItemRerollQualityStartopItemRerollQualityCancelopInstallResourceStartopInstallResourceCancelopInstallSilveropBuildingFillNutritionopBuildingChangeRenovationStateopBuildingBuySkinopBuildingClaimopBuildingGiveupopBuildingNutritionSilverStorageDepositopBuildingNutritionSilverStorageWithdrawopBuildingNutritionSilverRewardSetopConstructionSiteCreateopPlaceableObjectPlaceopPlaceableObjectPlaceCancelopPlaceableObjectPickupopFurnitureObjectUseopFarmableHarvestopFarmableFinishGrownItemopFarmableDestroyopFarmableGetProductopTearDownConstructionSiteopCastleGateUseopAuctionCreateRequestopAuctionCreateOfferopAuctionGetOffersopAuctionGetRequestsopAuctionBuyOfferopAuctionAbortAuctionopAuctionModifyAuctionopAuctionAbortOfferopAuctionAbortRequestopAuctionSellRequestopAuctionGetFinishedAuctionsopAuctionGetFinishedAuctionsCountopAuctionFetchAuctionopAuctionGetMyOpenOffersopAuctionGetMyOpenRequestsopAuctionGetMyOpenAuctionsopAuctionGetItemAverageStatsopAuctionGetItemAverageValueopContainerOpenopContainerCloseopContainerManageSubContaineropRespawnopSuicideopJoinGuildopLeaveGuildopCreateGuildopInviteToGuildopDeclineGuildInvitationopKickFromGuildopDuellingChallengePlayeropDuellingAcceptChallengeopDuellingDenyChallengeopChangeClusterTaxopClaimTerritoryopGiveUpTerritoryopChangeTerritoryAccessRightsopGetMonolithInfoopGetClaimInfoopGetAttackInfoopGetTerritorySeasonPointsopGetAttackScheduleopScheduleAttackopGetMatchesopGetMatchDetailsopJoinMatchopLeaveMatchopChangeChatSettingsopLogoutStartopLogoutCancelopClaimOrbStartopClaimOrbCancelopMatchLootChestOpeningStartopMatchLootChestOpeningCancelopDepositToGuildAccountopWithdrawalFromAccountopChangeGuildPayUpkeepFlagopChangeGuildTaxopGetMyTerritoriesopMorganaCommandopGetServerInfoopInviteMercenaryToMatchopSubscribeToClusteropAnswerMercenaryInvitationopGetCharacterEquipmentopGetCharacterSteamAchievementsopGetCharacterStatsopGetKillHistoryDetailsopLearnMasteryLevelopReSpecAchievementopChangeAvataropGetRankingsopGetRankopGetGvgSeasonRankingsopGetGvgSeasonRankopGetGvgSeasonHistoryRankingsopGetGvgSeasonGuildMemberHistoryopKickFromGvGMatchopGetChestLogsopGetAccessRightLogsopGetGuildAccountLogsopGetGuildAccountLogsLargeAmountopInviteToPlayerTradeopPlayerTradeCancelopPlayerTradeInvitationAcceptopPlayerTradeAddItemopPlayerTradeRemoveItemopPlayerTradeAcceptTradeopPlayerTradeSetSilverOrGoldopSendMiniMapPingopStuckopBuyRealEstateopClaimRealEstateopGiveUpRealEstateopChangeRealEstateOutlineopGetMailInfosopGetMailCountopReadMailopSendNewMailopDeleteMailopMarkMailUnreadopClaimAttachmentFromMailopUpdateLfgInfoopGetLfgInfosopGetMyGuildLfgInfoopGetLfgDescriptionTextopLfgApplyToGuildopAnswerLfgGuildApplicationopRegisterChatPeeropSendChatMessageopJoinChatChannelopLeaveChatChannelopSendWhisperMessageopSayopPlayEmoteopStopEmoteopGetClusterMapInfoopAccessRightsChangeSettingsopMountopMountCancelopBuyJourneyopSetSaleStatusForEstateopResolveGuildOrPlayerNameopGetRespawnInfosopMakeHomeopLeaveHomeopResurrectionReplyopAllianceCreateopAllianceDisbandopAllianceGetMemberInfosopAllianceInviteopAllianceAnswerInvitationopAllianceCancelInvitationopAllianceKickGuildopAllianceLeaveopAllianceChangeGoldPaymentFlagopAllianceGetDetailInfoopGetIslandInfosopAbandonMyIslandopBuyMyIslandopBuyGuildIslandopAbandonGuildIslandopUpgradeMyIslandopUpgradeGuildIslandopMoveMyIslandopMoveGuildIslandopTerritoryFillNutritionopTeleportBackopPartyInvitePlayeropPartyAnswerInvitationopPartyLeaveopPartyKickPlayeropPartyMakeLeaderopPartyChangeLootSettingopPartyMarkObjectopPartySetRoleopGetGuildMOTDopSetGuildMOTDopExitEnterStartopExitEnterCancelopQuestGiverRequestopGoldMarketGetBuyOfferopGoldMarketGetBuyOfferFromSilveropGoldMarketGetSellOfferopGoldMarketGetSellOfferFromSilveropGoldMarketBuyGoldopGoldMarketSellGoldopGoldMarketCreateSellOrderopGoldMarketCreateBuyOrderopGoldMarketGetInfosopGoldMarketCancelOrderopUnknown244opUnknown245opUnknown246opGoldMarketGetAverageInfoopSiegeCampClaimStartopSiegeCampClaimCancelopTreasureChestUsingStartopTreasureChestUsingCancelopUseLootChestopUseShrineopLaborerStartJobopLaborerTakeJobLootopLaborerDismissopLaborerMoveopLaborerBuyItemopLaborerUpgradeopBuyPremiumopBuyTrialopRealEstateGetAuctionDataopRealEstateBidOnAuctionopGetSiegeCampCooldownopFriendInviteopFriendAnswerInvitationopFriendCancelnvitationopFriendRemoveopInventoryStackopInventorySortopEquipmentItemChangeSpellopExpeditionRegisteropExpeditionRegisterCancelopJoinExpeditionopDeclineExpeditionInvitationopVoteStartopVoteDoVoteopRatingDoRateopEnteringExpeditionStartopEnteringExpeditionCancelopActivateExpeditionCheckPointopArenaRegisteropArenaRegisterCancelopArenaLeaveopJoinArenaMatchopDeclineArenaInvitationopEnteringArenaStartopEnteringArenaCancelopArenaCustomMatchopArenaCustomMatchCreateopUpdateCharacterStatementopBoostFarmableopGetStrikeHistoryopUseFunctionopUsePortalEntranceopResetPortalBindingopQueryPortalBindingopClaimPaymentTransactionopChangeUseFlagopClientPerformanceStatsopExtendedHardwareStatsopClientLowMemoryWarningopTerritoryClaimStartopTerritoryClaimCancelopRequestAppStoreProductsopVerifyProductPurchaseopQueryGuildPlayerStatsopQueryAllianceGuildStatsopTrackAchievementsopSetAchievementsAutoLearnopDepositItemToGuildCurrencyopWithdrawalItemFromGuildCurrencyopAuctionSellSpecificItemRequestopFishingStartopFishingCastingopFishingCastopFishingCatchopFishingPullopFishingGiveLineopFishingFinishopFishingCancelopCreateGuildAccessTagopDeleteGuildAccessTagopRenameGuildAccessTagopFlagGuildAccessTagGuildPermissionopAssignGuildAccessTagopRemoveGuildAccessTagFromPlayeropModifyGuildAccessTagEditorsopRequestPublicAccessTagsopChangeAccessTagPublicFlagopUpdateGuildAccessTagopSteamStartMicrotransactionopSteamFinishMicrotransactionopSteamIdHasActiveAccountopCheckEmailAccountStateopLinkAccountToSteamIdopBuyGvgSeasonBoosteropChangeFlaggingPrepareopOverChargeopOverChargeEndopRequestTrustedopChangeGuildLogoopPartyFinderRegisterForUpdatesopPartyFinderUnregisterForUpdatesopPartyFinderEnlistNewPartySearchopPartyFinderDeletePartySearchopPartyFinderChangePartySearchopPartyFinderChangeRoleopPartyFinderApplyForGroupopPartyFinderAcceptOrDeclineApplyForGroupopPartyFinderGetEquipmentSnapshotopPartyFinderRegisterApplicantsopPartyFinderUnregisterApplicantsopPartyFinderFulltextSearchopPartyFinderRequestEquipmentSnapshotopGetPersonalSeasonTrackerDataopUseConsumableFromInventoryopClaimPersonalSeasonRewardopEasyAntiCheatMessageToServeropSetNextTutorialStateopAddPlayerToMuteListopRemovePlayerFromMuteListopProductShopUserEventopGetVanityUnlocksopBuyVanityUnlocksopGetMountSkinsopSetMountSkinopSetWardrobeopChangeCustomizationopSetFavoriteIslandopGetGuildChallengePointsopTravelToHideoutopSmartQueueJoinopSmartQueueLeaveopSmartQueueSelectSpawnClusteropUpgradeHideoutopInitHideoutAttackStartopInitHideoutAttackCancelopHideoutFillNutritionopHideoutGetInfoopHideoutGetOwnerInfoopHideoutSetTributeopOpenWorldAttackScheduleStartopOpenWorldAttackScheduleCancelopOpenWorldAttackConquerStartopOpenWorldAttackConquerCancelopGetOpenWorldAttackDetailsopGetNextOpenWorldAttackScheduleTimeopRecoverVaultFromHideoutopGetGuildEnergyDrainInfoopChannelingUpdate"

var _OperationType_index = [...]uint16{0, 8, 14, 20, 35, 42, 56, 72, 86, 104, 121, 138, 155, 170, 194, 217, 237, 257, 283, 300, 323, 348, 354, 365, 377, 399, 417, 438, 460, 479, 501, 527, 548, 573, 589, 604, 620, 633, 652, 670, 692, 721, 745, 775, 800, 830, 844, 859, 871, 894, 918, 942, 967, 989, 1012, 1027, 1050, 1081, 1098, 1113, 1129, 1168, 1208, 1242, 1266, 1288, 1316, 1339, 1359, 1376, 1401, 1418, 1438, 1464, 1479, 1501, 1521, 1539, 1559, 1576, 1597, 1619, 1638, 1659, 1679, 1707, 1740, 1761, 1785, 1811, 1837, 1865, 1893, 1908, 1924, 1953, 1962, 1971, 1982, 1994, 2007, 2022, 2046, 2061, 2086, 2111, 2134, 2152, 2168, 2185, 2214, 2231, 2245, 2260, 2286, 2305, 2321, 2333, 2350, 2361, 2373, 2393, 2406, 2420, 2435, 2451, 2479, 2508, 2531, 2554, 2580, 2596, 2614, 2630, 2645, 2669, 2689, 2716, 2739, 2770, 2789, 2812, 2831, 2850, 2864, 2877, 2886, 2908, 2926, 2955, 2987, 3005, 3019, 3039, 3060, 3092, 3113, 3132, 3161, 3181, 3204, 3228, 3256, 3273, 3280, 3295, 3312, 3330, 3355, 3369, 3383, 3393, 3406, 3418, 3434, 3459, 3474, 3487, 3506, 3529, 3546, 3573, 3591, 3608, 3625, 3643, 3663, 3668, 3679, 3690, 3709, 3737, 3744, 3757, 3769, 3793, 3819, 3836, 3846, 3857, 3876, 3892, 3909, 3933, 3949, 3975, 4001, 4020, 4035, 4066, 4089, 4105, 4122, 4135, 4151, 4171, 4188, 4208, 4222, 4239, 4263, 4277, 4296, 4319, 4331, 4348, 4365, 4389, 4406, 4420, 4434, 4448, 4464, 4481, 4500, 4523, 4556, 4580, 4614, 4633, 4653, 4680, 4706, 4726, 4749, 4761, 4773, 4785, 4811, 4832, 4854, 4879, 4905, 4919, 4930, 4947, 4967, 4983, 4996, 5012, 5028, 5040, 5050, 5076, 5100, 5122, 5136, 5160, 5183, 5197, 5213, 5228, 5254, 5274, 5300, 5316, 5345, 5356, 5368, 5382, 5407, 5433, 5463, 5478, 5499, 5511, 5527, 5551, 5571, 5592, 5610, 5634, 5660, 5675, 5693, 5706, 5725, 5745, 5765, 5790, 5805, 5829, 5852, 5876, 5897, 5919, 5944, 5967, 5990, 6015, 6034, 6060, 6088, 6121, 6153, 6167, 6183, 6196, 6210, 6223, 6240, 6255, 6270, 6292, 6314, 6336, 6371, 6393, 6425, 6454, 6479, 6506, 6528, 6556, 6585, 6610, 6634, 6656, 6677, 6700, 6712, 6727, 6743, 6760, 6791, 6824, 6857, 6887, 6917, 6940, 6966, 7007, 7040, 7071, 7104, 7131, 7168, 7198, 7226, 7253, 7283, 7305, 7326, 7352, 7374, 7392, 7410, 7425, 7439, 7452, 7473, 7492, 7517, 7534, 7550, 7567, 7597, 7613, 7637, 7662, 7684, 7700, 7721, 7740, 7770, 7801, 7830, 7860, 7887, 7923, 7948, 7973, 7991}

func (i OperationType) String() string {
	if i >= OperationType(len(_OperationType_index)-1) {
		return "OperationType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OperationType_name[_OperationType_index[i]:_OperationType_index[i+1]]
}
