package operations

import "log"

func subtractionCallback(oldValue, newAddition interface{}) interface{} {
  switch t := oldValue.(type) {
   case int:
      switch newAddition.(type) {
        case int:  
		return oldValue.(int)-newAddition.(int)
        case float64:
          return float64(oldValue.(int)) - newAddition.(float64)
        default:
          log.Fatalf("invalid type '%T' on subtraction callback.", t)
          return oldValue
      }
    case float64:
      switch newAddition.(type) {
        case int:
          return oldValue.(float64) - float64(newAddition.(int))
        case float64:
          return oldValue.(float64) - newAddition.(float64)
        default:
          log.Fatalf("invalid type '%T' on subtraction callback.", t)
          return oldValue
      }

    default:
          log.Fatalf("invalid type '%T' on subtraction callback.", t)
          return 0
  }
  return oldValue	
}
