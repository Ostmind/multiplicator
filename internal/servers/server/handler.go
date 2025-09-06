package server

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"math/rand"
	"net/http"
)

type numberController struct {
	rtp    float64
	logger *slog.Logger
}

func newNumberHandler(rtp float64, log *slog.Logger) *numberController {
	return &numberController{rtp, log}
}

func (ctr numberController) generateNumberHandler(echo echo.Context) error {
	ctr.logger.Debug("Get Request for Numbers")

	response := map[string]interface{}{
		"result": generateNumber(ctr.rtp),
	}

	return echo.JSON(http.StatusOK, response)
}

func generateNumber(rtp float64) float64 {
	randVal := rand.Float64() // генерируем число из промежутка от 0 до 1

	chanceForMinVal := 1 - rtp                    // определяем вероятность числа 1
	chanceForMaxVal := rtp / 10000.0              // определяем вероятность числа 10000
	chanceForIntervalVal := rtp - chanceForMaxVal // Вероятность остальных чисел из промежутка

	//краевой случай когда надо отослать 1
	if randVal < chanceForMinVal {
		return 1.0
	}

	//краевой случай когда надо отослать 10000
	if randVal > 1-chanceForMaxVal {
		return 10000.0
	}

	//нормализуем вероятность
	p := (randVal - chanceForMinVal) / chanceForIntervalVal

	/*
		так как плотность у нас задана как f(t) = rtp / t^2
		проинтегрировав ее получаем F(t) = 1 - rtp / t
		следовательно обратная функция распределения для нормализованного p имеет вид :
		1 / (1 - p * F(Max при rtp = 1))
	*/

	multiplierVal := 1 / (1 - p*(1-1/10000))

	//если вдруг multiplierVal пробил 10000, тогда ограничиваем его 10000
	if multiplierVal > 10000 {
		multiplierVal = 10000
	}

	return multiplierVal
}
