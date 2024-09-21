package handler

import (
	"context"
	"errors"
	"fmt"
	"role/dao"
	"role/models"
	pb "role/proto/role"
	"strconv"
)

type Role struct{}

func (e *Role) AddRole(ctx context.Context, req *pb.AddRoleRequest, res *pb.AddRoleResponse) error {
	fmt.Println("AddRole---------", req)
	role := models.Role{}
	role.Title = req.GetTitle()
	role.Description = req.GetDescription()
	role.Status = 1

	if err := dao.DB.Create(&role).Error; err != nil {
		res.Status = -1
		res.Message = "增加角色失败，请重试"
		return err
	}
	res.Status = 0
	res.Message = "增加角色成功"
	return nil
}

func (e *Role) EditRole(ctx context.Context, req *pb.EditRoleRequest, res *pb.EditRoleResponse) error {
	fmt.Println("EditRole---------", req)
	role := models.Role{}
	role.ID = uint(req.GetId())

	var cnt int64
	if err := dao.DB.Find(&role).Count(&cnt).Error; err != nil {
		res.Status = -1
		res.Message = "修改角色信息失败，请稍后重试"
		return err
	}

	if cnt == 0 {
		res.Status = -1
		res.Message = "此角色不存在"
		return errors.New("此角色不存在")
	}

	if err := dao.DB.Model(&role).Updates(map[string]interface{}{"title": req.GetTitle(), "description": req.GetDescription()}).Error; err != nil {
		res.Status = -1
		res.Message = "修改角色信息失败，请稍后重试"
		return err
	}
	res.Status = 0
	res.Message = "修改角色信息成功"
	return nil
}

func (e *Role) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest, res *pb.DeleteRoleResponse) error {
	fmt.Println("DeleteRole---------", req)
	role := models.Role{}
	role.ID = uint(req.GetId())

	var cnt int64
	if err := dao.DB.Model(&models.Role{}).Where("id = ?", strconv.Itoa(int(req.GetId()))).Count(&cnt).Error; err != nil {
		res.Status = -1
		res.Message = "删除角色信息失败，请稍后重试"
		return err
	}

	if cnt == 0 {
		res.Status = -1
		res.Message = "此角色不存在"
		fmt.Println("没有了-================================")
		return errors.New("此角色不存在")
	}

	if err := dao.DB.Delete(&role).Error; err != nil {
		res.Status = -1
		res.Message = "删除角色失败，请稍后重试"
		return err
	}
	res.Status = 0
	res.Message = "删除角色成功"
	return nil
}

func (e *Role) GetRoleList(ctx context.Context, req *pb.GetRoleListRequest, res *pb.GetRoleListResponse) error {
	fmt.Println("GetRoleList---------", req)
	roleList := []models.Role{}
	if err := dao.DB.Find(&roleList).Error; err != nil {
		res.Status = -1
		res.Message = "获取角色列表失败，请稍后再试"
		return err
	}

	var temp []*pb.RoleItem
	for _, v := range roleList {
		temp = append(temp, &pb.RoleItem{
			Id:          int32(v.ID),
			Title:       v.Title,
			Description: v.Description,
			Status:      int32(v.Status),
		})
	}

	res.Status = 0
	res.Message = "获取角色列表成功"
	res.RoleList = temp
	return nil
}

func (e *Role) GetRoleInfo(ctx context.Context, req *pb.GetRoleInfoRequest, res *pb.GetRoleInfoResponse) error {
	fmt.Println("GetRoleInfo---------", req)
	role := models.Role{}
	role.ID = uint(req.GetId())
	if err := dao.DB.Find(&role).Error; err != nil {
		res.Status = -1
		res.Message = "获取角色信息失败，请稍后重试"
		return err
	}
	res.Status = 0
	res.Message = "获取角色信息成功"
	res.RoleInfo = &pb.RoleItem{
		Id:          int32(role.ID),
		Title:       role.Title,
		Description: role.Description,
		Status:      int32(role.Status),
	}
	return nil
}
