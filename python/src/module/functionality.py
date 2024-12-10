GLOBAL_VAR = "I am changed global variable"

hell, nah = "hell", "nah" 
# new comment
# one more



def global_function():
    print("This is a global function.")
    print(GLOBAL_VAR)
    print(f"{hell} {nah}")


class MyClass:
    class_variable = "I am changed class variable"

    def __init__(self, instance_variable):
        self.instance_variable = instance_variable

    def display_instance_variable(self):
        function_variable = "I am new function scope variable"
        print(f"Instance Variable Updated: {function_variable}")

    @classmethod
    def display_class_variable(cls):
        def one_more_change():
            print("+++")
        print(f"Class Variable: {cls.class_variable}")

    @staticmethod
    def static_method_example():
        print("This is a static method, it doesn't access instance or class variables.")


def function_with_local_variable():
    local_variable = "I am a local variable"
    print(local_variable)

def new_global_function():
    new_variable = "Im new global function variable"
    print(new_variable)