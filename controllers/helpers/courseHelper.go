package helpers

import (
	"github.com/google/uuid"
	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/models/requests"
	"github.com/meja_belajar/models/responses"
	"github.com/meja_belajar/utils"
)


func GetCourse(GetCourseRequestDTO requests.GetCourseRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var course database.Courses

	err := db.First(&course, utils.StringToUUID(GetCourseRequestDTO.CourseID)).Error
	if err != nil {
		return 500, err.Error()
	}

	if course.ID == uuid.Nil {
		return 404, "Course not found"
	}

	output := outputs.GetCourseOutput{}
	output.Code = 200
	output.Message = "Success"
	output.Data = responses.CourseResponseDTO{
		CourseID: 	  	 course.ID.String(),
		Name:     		 course.Name,
		Detail:   		 course.Detail,
		Rating:   		 course.Rating,
		HourlyRate: 	 course.HourlyRate,
		CourseStartTime: course.CourseStartTime,
		CourseEndTime: 	 course.CourseEndTime,
	}
	return 200, output
}

func AddCourse(AddCourseRequestDTO requests.AddCourseRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	course := database.Courses{
		Name			:    AddCourseRequestDTO.Name,
		Detail			:    AddCourseRequestDTO.Detail,
		Rating			:    AddCourseRequestDTO.Rating,
		HourlyRate		: 	 AddCourseRequestDTO.HourlyRate,
		CourseStartTime	: 	 AddCourseRequestDTO.CourseStartTime,
		CourseEndTime	: 	 AddCourseRequestDTO.CourseEndTime,
		IsActive		:  	 AddCourseRequestDTO.IsActive,
		CreatedBy		: 	 AddCourseRequestDTO.CreatedBy,
	}

	err := db.Create(&course).Error
	if err != nil {
		return 500, err.Error()
	}

	output := outputs.AddCourseOutput{}
	output.Code = 200
	output.Message = "Success"
	output.Data = responses.CourseResponseDTO{
		CourseID		:   course.ID.String(),
		Name			:   course.Name,
		Detail			:   course.Detail,
		Rating			:   course.Rating,
		HourlyRate		: 	course.HourlyRate,
		CourseStartTime	: 	course.CourseStartTime,
		CourseEndTime	: 	course.CourseEndTime,
	}
	return 200, output
}

func UpdateCourse(UpdateCourseRequestDTO requests.UpdateCourseRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var course database.Courses

	err := db.First(&course, "id = ?", utils.StringToUUID(UpdateCourseRequestDTO.CourseID)).Error
	if err != nil {
		return 500, err.Error()
	}

	if course.ID == uuid.Nil {
		return 404, "Course not found"
	}

	course.Name = UpdateCourseRequestDTO.Name
	course.Detail = UpdateCourseRequestDTO.Detail
	course.Rating = UpdateCourseRequestDTO.Rating
	course.HourlyRate = UpdateCourseRequestDTO.HourlyRate
	course.CourseStartTime = UpdateCourseRequestDTO.CourseStartTime
	course.CourseEndTime = UpdateCourseRequestDTO.CourseEndTime
	course.IsActive = UpdateCourseRequestDTO.IsActive
	course.UpdatedBy = UpdateCourseRequestDTO.UpdatedBy

	err = db.Save(&course).Error
	if err != nil {
		return 500, err.Error()
	}

	output := outputs.UpdateCourseOutput{}
	output.Code = 200
	output.Message = "Success"
	output.Data = responses.CourseResponseDTO{
		CourseID		:   course.ID.String(),
		Name			:   course.Name,
		Detail			:   course.Detail,
		Rating			:   course.Rating,
		HourlyRate		: 	course.HourlyRate,
		CourseStartTime	: 	course.CourseStartTime,
		CourseEndTime	: 	course.CourseEndTime,
	}
	return 200, output
}