GLOBAL_VAR = "I am changed global variable"


def global_function():
    print("This is a global function.")
    print(GLOBAL_VAR)


class MyClass:
    class_variable = "I am a class variable"

    def __init__(self, instance_variable):
        self.instance_variable = instance_variable

    def display_instance_variable(self):
        print(f"Instance Variable Updated: {self.instance_variable}")

    @classmethod
    def display_class_variable(cls):
        print(f"Class Variable: {cls.class_variable}")

    @staticmethod
    def static_method_example():
        print("This is a static method, it doesn't access instance or class variables.")


def function_with_local_variable():
    local_variable = "I am a local variable"
    print(local_variable)
